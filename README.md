This simple tool divides the recent system load by the number of processors to get a meaningfull load number for servers with different numbers of processors.

I needs read access to /proc/loadavg and /proc/cpuinfo and works only for systems with these "files".

Usually I would implement a functionality like this in a shell script. But it turned out that `expr` doesn't work with floats and `bc` isn't installed by default on some systems. So - for my needs - downloading the binary provided in this package is less work than installing `bc` and deploying a shell script.

If anybody else has this need: voilà. If anybody knows a better way: Please contact me or open an issue. If anybody knows how to imporove my - naive - go code: Please let me known and open an issue.

Get the downloads on the Releases page: https://github.com/niko/sysload/releases
