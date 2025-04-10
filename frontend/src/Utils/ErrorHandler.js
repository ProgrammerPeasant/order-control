export const handleErrorMessage = (error, isLoginPage = false) => {
    if (error.response) {
        const {status, data} = error.response;
        const errorMessages = {
            400: "Invalid data or incorrect request",
            401: isLoginPage
                ? "Incorrect login or password"
                : "User is not authorized",
            403: "You do not have sufficient permissions to perform this action",
            404: "The requested object was not found or does not exist",
            500: "Server error, please try again later"
        };

        return error.response.data.error || errorMessages[status] || `Unknown error: (${data.message || "Something went wrong"})`;
    } else if (error.request) {
        return "No response from the server, please check your internet connection"
    } else {
        return `Request error: ${error.message}`;
    }
};