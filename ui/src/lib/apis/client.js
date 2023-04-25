import wretch from "wretch";
import Cookies from "js-cookie";

const dev = false;

const csrfMiddleware = next => (url, opts) => {
	const token = Cookies.get("_csrf");
	if (token) opts.headers = {...opts.headers, "X-CSRF-Token": token};
	return next(url, opts);
};

export const createApiClient = (base) => {
	return wretch("/api" + base).middlewares([csrfMiddleware])
}