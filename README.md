# kube-search

[![Build Status](https://travis-ci.org/mauricioklein/kube-search.svg?branch=master)](https://travis-ci.org/mauricioklein/kube-search)
![GitHub release](https://img.shields.io/github/release/mauricioklein/kube-search.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Maintainability](https://api.codeclimate.com/v1/badges/5369e243c130ce10f2bc/maintainability)](https://codeclimate.com/github/mauricioklein/kube-search/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/5369e243c130ce10f2bc/test_coverage)](https://codeclimate.com/github/mauricioklein/kube-search/test_coverage)

Fuzzy search K8s fields path by namespace and resource name

## Requirements

- Go 1.11 or superior
- Access to [kubectl](https://github.com/kubernetes/kubectl)

## Instalation

```bash
$ go get -v github.com/mauricioklein/kube-search
```

## How it works?

`kube-search` doesn't have any direct dependency to [kubectl official module](https://github.com/kubernetes/kubectl). Instead, the tool relies on the kubectl client available on the running environment. Thus, make
sure you have your `kubectl` setup and available in `$PATH`, so `kube-search` can interact with it. New releases
of `kubectl` doesn't require any change on `kube-search`, as long the interface is backward compatible.

## Usage

`kube-search` requires two arguments:
- the root namespace (e.g. "po.spec")
- the resource name (e.g. "livenessProbe")

`kube-search` returns the single resource with the highest matching score among all the resources under the provided namespace.

```bash
$ kube-search -n po.spec -r livenessProbe
# po.spec.containers.livenessProbe
```

The matching score calculation is performed using [Jaro distance](https://rosettacode.org/wiki/Jaro_distance) algorithm, provided by the library [smetrics](https://github.com/xrash/smetrics). Thus, while the matching might not be exact (e.g. a typo on the resource name), `kube-search` is still capable of returning an accurate result.

To return a larger set of results, provide the flag `-c`, followed by the number of results desired:

```bash
$ kube-search -n po.spec -r livenessProbe -c 3
# po.spec.containers.livenessProbe
# po.spec.initContainers.livenessProbe
# po.spec.initContainers.readinessProbe
```

The `show-score` flag can be used to display the matching score along with the results:

```bash
$ kube-search -n po.spec -r livenessProbe -c 3
# po.spec.containers.livenessProbe (matching score: 1.000000)
# po.spec.initContainers.livenessProbe (matching score: 1.000000)
# po.spec.initContainers.readinessProbe (matching score: 0.725774)
```

Finally, for all the available options, please refer to the help section:

```bash
$ kube-search -h
```

## Testing

```bash
$ go test -v ./...
```

## Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request

## License

`kube-search` is licensed under the [MIT](https://opensource.org/licenses/MIT) License.
