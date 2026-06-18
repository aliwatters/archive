function uri_lowercase(r) {
	return r.uri.toLowerCase().replace(/\n|\r/g, "");
}

export default { uri_lowercase };
