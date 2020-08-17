[drone]: https://cloud.drone.io
[drone repo]: https://cloud.drone.io/magicalbanana/peggy-hill
[drone badge master]: https://cloud.drone.io/api/badges/magicalbanana/peggy-hill/status.svg
[doc]: https://godoc.org/github.com/magicalbanana/peggy-hill
[doc badge]: https://godoc.org/github.com/magicalbanana/peggy-hill?status.svg
[go report card]: https://goreportcard.com/report/github.com/magicalbanana/peggy-hill
[go report card badge]: https://goreportcard.com/badge/github.com/magicalbanana/peggy-hill

| master                                                    |
| -                                                         |
| [![drone repo][drone badge master]][drone repo]           |
| [![doc][doc badge]][doc]                                  |
| [![go report card][go report card badge]][go report card] |

# Peggy Hill

> Peggy: What if I'm not really as smart as I think I am?!

> Hank: aw c'mon Peggy, you've got an IQ of 170, you said so lots of times

> Peggy: yes but there could be a margin of error, especially when it's just my own estimate

## Author's Notes

I spent about a good half an hour before realizing that the math behind this
problem is beyond my current knowledge. So like any good engineer, I Googled
the formula for calculating the gear sizes and discovered the _secret society_
of the **Google Foo Bar** challenges.

I want to acknowledge before hand that the code I have written is not solely
thought on my own but was insipired by theories and implementations from
engineers that have attempted the challenge.

I do want to emphasize that I wanted to make sure that I did gain something
from this challenge so I dug deeper on the theories behind the implementations
of the engineers and also understand what exactly a gear train[^1] is to
better understand my solution and be able to write it in different languages.

There were quite a few solutions but I mostly focused on this explanation[^2]
and this one[^3] as it gave me a better explanation behind the math which
allowed me to be able to implement the code I need to provide a solution in
the languages of my choosing (apart from the required language, elixir).

## Solutions

All the applications are containerized for ease of running locally. There are
two ways to do this:

1. Run the containers locally by building them using the `docker-compose.yml`
   config.
2. Run the images locally by pulling the images from [hub.docker.com](https://hub.docker.com).
 - [svajone/peggo:latest](https://hub.docker.com/repository/docker/svajone/peggy-hill-go)
 - [svajone/peggelixir:latest](https://hub.docker.com/repository/docker/svajone/peggy-hill-elixir)

### Go

#### Usage

The binary for the `go` container is called `peggo` because I like word play.


Through docker-compose:
```
$ docker-compose run --rm peggo /bin/sh
$ cat testdata/test_data_1.txt | peggo
```

Directly in the image:
```
docker run -it --rm --mount src="$(pwd)/testdata",target="/app/testdata",type=bind svajone/peggo:latest
```

### Elixir

## References

- https://www.mathsisfun.com/algebra/systems-linear-equations-matrices.html
- https://gist.github.com/jlacar/e66cd25bac47fab887bffbc5e924c2f5
- https://jasonqsy.github.io/algorithm/2018/01/06/Google-foobar.html

[^1]: https://en.wikipedia.org/wiki/Gear_train
[^2]: https://gist.github.com/jlacar/e66cd25bac47fab887bffbc5e924c2f5
[^3]: https://jasonqsy.github.io/algorithm/2018/01/06/Google-foobar.html
