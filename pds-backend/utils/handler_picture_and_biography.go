package utils

func HandlerPictureAndBiography(picture *string, bio *string) (string, string) {
	var resultPicture, resultBio string
	if picture != nil || bio != nil {
		resultPicture = *picture
		resultBio = *bio
		return resultPicture, resultBio
	} else {
		return "", ""
	}
}