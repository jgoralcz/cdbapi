package images

// ImageByID is the query to find the metadata on an image by ID
var ImageByID = `
	SELECT image_id, waifu_id, image_url_path_extra, image_url_clean_path_extra,
	nsfw, width, height, file_type, buffer_length, buffer_length_clean
	FROM waifu_schema.waifu_table_images
	WHERE image_id = $1;
`
