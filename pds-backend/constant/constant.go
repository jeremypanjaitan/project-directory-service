package constant

const YES = "Y"
const NO = "N"
const UUID = "uuid"
const USERNAME = "username"
const EMAIL = "Email"
const BEARER = "Bearer "
const ACCESS_UUID = "AccessUUID"
const SUCCESS = "Success"

const FIREBASE_EMAIL_VER = "https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key=%s"
const GOOGLE_VERIFY_PASSWORD = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=%s"
const GOOGLE_SIGN_IN_PASSWORD = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s"
const GOOGLE_VERIFY_CUSTOM_TOKEN = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=%s"
const GOOGLE_REFRESH_TOKEN = "https://securetoken.googleapis.com/v1/token?key=%s"
const VERIFY_EMAIL = "VERIFY_EMAIL"
const USERS_COLLECTION = "users"
const IMAGE_NOTIFICATION = "https://firebasestorage.googleapis.com/v0/b/project-directory-servic-a843a.appspot.com/o/pictures%2FLogo.png?alt=media&token=437b0e02-1d71-47d9-a937-0d57a0fecb42"
const TITLE_NOTIFICATION = "ITFUN"

//email subject
const EMAIL_VER_SUB = "Email Verification"
const PASSWORD_CHANGE_SUB = "Change Password Link"

//activity constant
const LIKE = "LIKE"
const DISLIKE = "DISLIKE"
const COMMENTED = "COMMENTED"

func ActivityHeaderConstant(typesOf string, projectTitle string, projectCreator string) string {
	var result string

	switch typesOf {
	case LIKE:
		result = "You liked a project: " + projectTitle
	case DISLIKE:
		result = "You unliked a project: " + projectTitle
	case COMMENTED:
		result = "You commented on " + projectCreator + "'s project."
	}

	return result
}
func NotificationMessageConstant(typesOf string, projectTitle string, userName string) string {
	var result string

	switch typesOf {
	case LIKE:
		result = userName + " liked your project: " + projectTitle
	case COMMENTED:
		result = userName + " commented on your project : " + projectTitle
	}
	return result
}
