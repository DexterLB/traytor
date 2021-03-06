package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/DexterLB/mvm/progress"
	"github.com/DexterLB/traytor/hdrimage"
	"github.com/DexterLB/traytor/random"
	"github.com/DexterLB/traytor/raytracer"
	"github.com/DexterLB/traytor/rpc"
	"github.com/DexterLB/traytor/scene"
	"github.com/codegangsta/cli"
)

func renderer(
	width, height int,
	renderedImages chan *hdrimage.Image,
	scene *scene.Scene,
	seed int64,
	totalSamples int,
	threads int,
	quiet bool,
) {
	randomGen := random.New(seed)

	sampleCounter := rpc.NewSampleCounter(totalSamples)
	var bar *progress.ProgressBar
	if !quiet {
		bar = progress.StartProgressBar(totalSamples, "rendering samples")
	}

	wg := sync.WaitGroup{}
	wg.Add(threads)

	defer func() {
		wg.Wait()
		if !quiet {
			bar.Done()
		}
	}()

	for i := 0; i < threads; i++ {
		go func(seed int64) {
			defer wg.Done()

			raytracer := raytracer.Raytracer{
				Scene:  scene,
				Random: random.New(seed),
			}

			image := hdrimage.New(width, height)
			image.Divisor = 0

			for {
				if sampleCounter.Dec(1) == 0 {
					renderedImages <- image
					return
				}
				raytracer.Sample(image)
				if !quiet {
					bar.Add(1)
				}
			}
		}(randomGen.NewSeed())
	}
}

func runRender(c *cli.Context) error {
	scenePath, image := getArguments(c)
	quiet := c.GlobalBool("quiet")

	if !quiet {
		log.Printf(
			"will render %d samples of %s to %s of size %dx%d with %d threads",
			c.Int("total-samples"),
			scenePath, image,
			c.Int("width"), c.Int("height"),
			c.Int("max-jobs"),
		)
	}

	width, height := c.Int("width"), c.Int("height")
	totalSamples := c.Int("total-samples")
	threads := c.Int("max-jobs")

	renderedImages := make(chan *hdrimage.Image)

	scene, err := scene.LoadFromFile(scenePath)
	if err != nil {
		return fmt.Errorf("can't open scene: %s", err)
	}
	scene.Init()

	go func() {
		renderer(width, height, renderedImages, scene, 42, totalSamples, threads, quiet)
		close(renderedImages)
	}()

	averageImage := hdrimage.New(width, height)
	averageImage.Divisor = 0
	for currentImage := range renderedImages {
		averageImage.Add(currentImage)
		averageImage.Divisor += currentImage.Divisor
	}

	return saveImage(averageImage, image, c.String("format"))
}
