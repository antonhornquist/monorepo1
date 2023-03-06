Feature: Photo collection and consumption
	Background: Photos with metadata are available for upload
		Given a collection of photos in JPEG format
		And the JPEG files of the collection contain metadata describing aperture, shutter speed and ISO at point of shooting each photo
		And the JPEG files of the collection contain metadata describing when photo was taken

	Scenario: An uploaded photo is presented in the UI
		When a photo from the collection is uploaded
		Then it should become available in the UI

	Scenario: A collection of uploaded photos is presented together in the UI
		When multiple photos in the collection are uploaded
		Then they should become available in a reel in the UI
		And they should be sorted by when photo was taken

	Scenario: Uploaded photos should have its metadata scanned automatically
		When a photo from the collection is uploaded
		Then the aperture, shutter speed and ISO used when shooting the photo is presented for each photo in the UI
		And the date photo was taken is presented for each photo in the UI
