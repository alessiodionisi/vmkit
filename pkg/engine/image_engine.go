package engine

func (eng *Engine) FindImage(name string) *Image {
	image, exist := eng.images[name]
	if !exist {
		return nil
	}

	return image
}

func (eng *Engine) ListImages() []*Image {
	images := make([]*Image, 0, len(eng.images))
	for _, img := range eng.images {
		images = append(images, img)
	}

	return images
}
