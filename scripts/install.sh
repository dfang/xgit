#!/bin/sh
set -eu

info() {
    echo "$@" >&2
}

debug() {
    echo "$@" >&2
}

error() {
	echo "$@" >&2
	exit 1
}

get_os() {
	os="$(uname -s)"
	if [ "$os" = Darwin ]; then
		echo "Darwin"
	elif [ "$os" = Linux ]; then
		echo "Linux"
	else
		error "unsupported OS: $os"
	fi
}

get_arch() {
	arch="$(uname -m)"
	if [ "$arch" = x86_64 ]; then
		echo "x86_64"
	elif [ "$arch" = aarch64 ] || [ "$arch" = arm64 ]; then
		echo "arm64"
	else
		error "unsupported architecture: $arch"
	fi
}

download_file() {
	url="$1"
	filename="$(basename "$url")"
	# cache_dir="$(mktemp -d)"
	cache_dir="/tmp"
	file="$cache_dir/$filename"

	info "xgit: installing xgit..."

	if command -v curl >/dev/null 2>&1; then
		debug ">" curl -fLlSso "$file" "$url"
		curl -fLlSso "$file" "$url"
	else
		if command -v wget >/dev/null 2>&1; then
			debug ">" wget -qO "$file" "$url"
			stderr=$(mktemp)
			wget -O "$file" "$url" >"$stderr" 2>&1 || error "wget failed: $(cat "$stderr")"
		else
			error "xgit standalone install requires curl or wget but neither is installed. Aborting."
		fi
	fi

	echo "$file"
}

install_xgit() {
	# download the tarball
	version="v0.0.6"
	os="$(get_os)"
	arch="$(get_arch)"
	xdg_data_home="${XDG_DATA_HOME:-$HOME/bin}"
	install_path="${XGIT_INSTALL_PATH:-$xdg_data_home/xgit}"
	install_dir="$(dirname "$install_path")"
	tarball_url="https://ghproxy.com/https://github.com/dfang/xgit/releases/download/${version}/xgit_${os}_${arch}.tar.gz"
	cache_file=$(download_file "$tarball_url")
	# extract tarball
	mkdir -p "$install_dir"
	rm -rf "$install_path"
    export TMPDIR=/tmp
	cd "$(mktemp -d)"
	tar -xzf "$cache_file"
	mv xgit "$install_path"
	info "xgit: installed successfully to $install_path"
}

after_finish_help() {
	case "${SHELL:-}" in
	*)
		info "xgit: run \`$install_path --help\` to get started"
		;;
	esac
}

install_xgit
after_finish_help