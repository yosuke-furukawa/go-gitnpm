go-gitnpm
======================

npm wrapper for unpublished npm modules

npm could install library from some git repositories.
but npm user should set the library like this.

```shell
$ npm install git+ssh://git@github.com:yosuke-furukawa/gitnpm.git#master
```

this is so long....

*`gitnpm`* could make the command shorter.

```shell
$ go-gitnpm install yosuke-furukawa/gitnpm
// npm install git+ssh://git@github.com:yosuke-furukawa/gitnpm.git#master
```

```shell
$ go-gitnpm install yosuke-furukawa/gitnpm v0.1.0
// npm install git+ssh://git@github.com:yosuke-furukawa/gitnpm.git#v0.1.0
```

If you have github enterprise, you can set the path using setup command

```shell
$ go-gitnpm setup --hostname github.enterprise.com
$ go-gitnpm install yosuke-furukawa/gitnpm
// npm install git+ssh://git@github.enterprise.com:yosuke-furukawa/gitnpm.git#master
```

Install
======================
