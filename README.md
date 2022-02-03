# bump-version

 Bump a given semantic version, following a given version fragment.

## Install

### manually

Download the pre-compiled binaries from the [OSS releases page](https://github.com/wesleimp/bump-version/releases) and copy them to the desired location.

### deb and rpm packages

Download the `.deb` or `.rpm` packages from the [OSS releases page](https://github.com/wesleimp/bump-version/releases) and install them with the appropriate tools.

### compiling from source

If you just want to build from source, follow these steps:

**clone**

```sh
git clone https://github.com/wesleimp/bump-version
cd bump-version 
```

**dependencies**

```sh
go mod tidy
```

For the next steps, you can just run `make install` to build binaries directly inside `/usr/local/bin/` folder. Or just follow the steps below:

**build**

```sh
make build
```

**verify it works**

```sh 
./bin/bump-version -v
```

## Usage

```sh 
bump-version [options...] <version>
```

The available options for the `fragment` flag are `[major | feature | bug | alpha | beta | rc]`. See some examples:

| fragment         | version         |   | output        |
| ---------------- | --------------- | - | ------------- |
| major            | 2.11.7          |   | 3.0.0         |
| major            | 2.11.7-alpha3   |   | 3.0.0         |
| feature          | 2.11.7          |   | 2.12.0        |
| feature          | 2.11.7-alpha3   |   | 2.12.0        |
| bug              | 2.11.7          |   | 2.11.8        |
| bug              | 2.11.7-alpha3   |   | 2.11.8        |
| alpha            | 2.11.7          |   | 2.11.7-alpha1 |
| alpha            | 2.11.7-alpha3   |   | 2.11.7-alpha4 |
| beta             | 2.11.7          |   | 2.11.7-beta1  |
| beta             | 2.11.7-alpha3   |   | 2.11.7-beta1  |
| rc               | 2.11.7          |   | 2.11.7-rc1    |
| rc               | 2.11.7-alpha3   |   | 2.11.7-rc1    |

## LICENSE

[MIT](./LICENSE)
