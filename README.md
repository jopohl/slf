# slf - Show Large Files

Over time your brand-new disk gets fuller and fuller making you wonder _Where is all my disk space at?_. 
Meet ``` slf ``` which deeply scans a directory and shows the largest files or subdirectories in it. 
The tool supports Windows, Linux and macOS.

### Usage examples
#### Show 6 largest files overall
```bash
$ slf -n 6 /
90.23 GiB /var/lib/libvirt/images/win10.qcow2
90.01 GiB /var/lib/libvirt/images/ubuntu20.04.qcow2
80.01 GiB /var/lib/libvirt/images/win7.qcow2
50.01 GiB /var/lib/libvirt/images/MX_Linux.qcow2
```


#### Show 5 largest directories in /usr
```bash
$ slf -d -n 5 /var
15.94 GiB /usr/local
7.142 GiB /usr/lib
5.809 GiB /usr/share
1.535 GiB /usr/bin
499.8 MiB /usr/lib32
```

### Download

The easiest way to download the latest version of ```slf``` is to use the bundled ```download.py``` which automatically detects your OS and places the according version of ```slf``` into the current directory.
```bash
git clone https://github.com/jopohl/slf
cd slf
python download.py
./slf    # Running without any argument will scan the current directory
```

To save some typing on Unix shells, use the following oneliner.
```bash
python -c "$(curl -fsSL https://raw.githubusercontent.com/jopohl/slf/main/download.py)"
```

Alternatively, manually download the latest version of ```slf``` for your OS with one of the links below.
- [Linux](https://github.com/jopohl/slf/releases/latest/download/slf-linux-amd64.tar.gz)
- [Windows](https://github.com/jopohl/slf/releases/latest/download/slf-windows-amd64.zip)
- [macOS](https://github.com/jopohl/slf/releases/latest/download/slf-darwin-amd64.tar.gz)


### Why not simply use tool X, Y or Z?

My main motivation for this project was to implement something in Go. 
Therefore, I did not make an extensive study about existing tools. 
Nevertheless, ``` slf ``` is built with performance and concurrency in mind.

For example, listing the 20 largest files in ``` /var ``` with default Linux tools
```bash
sudo find /var -type f -printf "%s\t%p\n" | sort -n | tail -20
```
takes 5.585 seconds on my system<sup>[[1]](#mysystem)</sup>. Compared to that 
```bash
sudo slf -n 20 /var
```
takes only 1.071 seconds. Not too bad!

<a name="mysystem">[1]</a>: Intel i7-6700K, Samsung 970 Pro 1 TB NVMe M.2 SSD
