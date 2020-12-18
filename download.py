#!/usr/bin/env python

from __future__ import print_function

import os
import sys
import tarfile
import tempfile
import zipfile

try:
    # noinspection PyCompatibility
    from urllib.request import urlopen
except ImportError:  # Python 2 legacy
    # noinspection PyCompatibility,PyUnresolvedReferences
    from urllib2 import urlopen


def download_to_tmp(url):
    data = urlopen(url).read()

    suffix = ".tar.gz" if url.endswith(".tar.gz") else ".zip" if url.endswith(".zip") else ""
    fd, filename = tempfile.mkstemp(suffix=suffix)

    with open(filename, "wb") as f:
        f.write(data)
    os.close(fd)

    return filename


def extract(filename):
    if filename.endswith(".tar.gz"):
        tar = tarfile.open(filename, "r:gz")
        tar.extractall()
        tar.close()
    elif filename.endswith(".zip"):
        with zipfile.ZipFile(filename, 'r') as f:
            f.extractall()


if __name__ == '__main__':
    if sys.platform.startswith("linux"):
        url = "https://github.com/jopohl/slf/releases/latest/download/slf-linux-amd64.tar.gz"
    elif sys.platform.startswith("win32"):
        url = "https://github.com/jopohl/slf/releases/latest/download/slf-windows-amd64.zip"
    elif sys.platform.startswith("darwin"):
        url = "https://github.com/jopohl/slf/releases/latest/download/slf-darwin-amd64.tar.gz"
    else:
        print("OS {} not supported".format(sys.platform))
        sys.exit(1)

    fname = download_to_tmp(url)
    extract(fname)
    os.remove(fname)
