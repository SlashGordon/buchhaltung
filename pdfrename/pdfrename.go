package pdfrename

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"

	ty "github.com/SlashGordon/buchhaltung/types"
	"github.com/SlashGordon/buchhaltung/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
)

// Start ...
func Start(conf ty.RenameConfig, inputPath string, outputPath string) error {
	// pre processing
	// docker pull jbarlow83/ocrmypdf
	// docker run -v $PWD:/data --workdir /data --rm -i jbarlow83/ocrmypdf 201912151411.pdf 201912151411ocr.pdf
	ocrPath := path.Join(inputPath, "ocr")
	if !utils.DirExists(ocrPath) {
		os.MkdirAll(ocrPath, os.ModePerm)
	}
	pdfFiles, err := utils.WalkMatch(inputPath, "*.pdf")
	for _, pdfFile := range pdfFiles {
		if runInDocker("jbarlow83/ocrmypdf:latest", inputPath, pdfFile, path.Join("ocr", pdfFile)) {
			utils.Logger.Infof("OCR for file %v was successful.", pdfFile)
		} else {
			utils.Logger.Warnf("OCR for file %v was unsuccessful.", pdfFile)
		}
	}
	err = conf.Rename(inputPath, outputPath)
	return err
}

func runInDocker(image string, vol string, input string, output string) bool {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		utils.Logger.Warn("Couldn't start ocrmypdf because of missing docker dependencies.")
		return false
	}

	reader, err := cli.ImagePull(
		ctx,
		fmt.Sprintf("docker.io/%v", image),
		types.ImagePullOptions{})
	if err != nil {
		utils.Logger.Warn("Couldn't start ocrmypdf because of failed pull.")
		utils.Logger.Error(err)
		return false
	}

	io.Copy(os.Stdout, reader)

	resp, err := cli.ContainerCreate(ctx,
		&container.Config{
			Image:      image,
			WorkingDir: "/data",
			Cmd:        []string{input, output},
			Tty:        true,
		},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: vol,
					Target: "/data",
				},
			},
		}, nil, "")

	if err != nil {
		utils.Logger.Warn("Couldn't create container.")
		return false
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		utils.Logger.Warn("Couldn't start container.")
		return false
	}

	status, _ := cli.ContainerWait(ctx, resp.ID)

	if status != 0 {
		utils.Logger.Warnf("OCR for file %v failed with status %i", input, status)
		return false
	}
	return true
}
