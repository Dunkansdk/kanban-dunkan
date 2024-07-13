package footer

func (footer *Model) UpdateContent(mode, breadcrumb string) {
	footer.Mode = mode
	footer.Breadcrumb = breadcrumb
}
