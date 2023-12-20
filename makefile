download-depedencies:
	@go mod download

remove-images:
	@rm images/*.png

run-pokedex:
	@go run cmd/pokedex/app.go

run-data-scraper:
	@go run cmd/data_scraper/pokemon_scraper.go

run-ascii-image-generator:
	@go run cmd/ascii_images_generator/generator.go