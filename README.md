# sync-to-clickhouse

Sync OLTP database[^1] to ClickHouse.

[^1]: Such as MySQL from [canal](https://github.com/alibaba/canal).

## How to install and run

1. 首先运行 `make` 命令编译代码，编译完成后会在 [build](manifest/build) 目录下生成一个
   `sync-to-clickhouse.linux-amd64v3.tar.xz` 的压缩包。 解压缩后会得到一个 `sync-to-clickhouse` 的二进制文件。

2. 复制 [config.example](manifest/config.example) 目录到你的工作目录下，并重命名为 `config`。
   该目录下有一个配置文件 [config.toml](manifest/config.example/config.toml)。 你可以根据自己的需要修改这个配置文件。

3. 运行 `./sync-to-clickhouse install` 命令， 该命令会将 `sync-to-clickhouse` 注册到 systemd 服务中。

4. `systemctl start sync-to-clickhouse.service` 启动服务。

5. `systemctl enable sync-to-clickhouse.service` 设置开机自启。

## How to maintain

- 更新二进制文件：将 `make` 之后的压缩包移动到工作目录，使用 [auto.sh](manifest/deploy/auto.sh) 脚本一键部署。

- 卸载服务，使用 `./sync-to-clickhouse uninstall` 命令。

- 查看日志，使用 `journalctl -u sync-to-clickhouse.service` 命令。

- 查看服务状态，使用 `systemctl status sync-to-clickhouse.service` 命令。

## How to full sync from mysql

1. 按需修改 [sync.sql](manifest/maintain/sync.sql)，并同时将 [full-sync.sh](manifest/maintain/full-sync.sh) 复制到工作目录下。

2. 修改 [full-sync.sh](manifest/maintain/full-sync.sh) 中的配置。

3. 运行 `./full-sync.sh` 命令，开始全量同步。

### sync new table from mysql

同样使用 [sync.sql](manifest/maintain/sync.sql) 来同步新表。

## MySQL to ClickHouse table structure conversion

GPT Prompt

```text
将MySQL的建表语句改成ClickHouse的建表语句：
Engine采用ReplacingMergeTree()。
ORDER BY来自于MySQL表的主键（例如`ORDER BY id`）。
ClickHouse不支持`NOT NULL`语句。
其他类型的属性根据MySQL建表语句确定，凡是DEFAULT NULL的类型，都要加上Nullable()。
凡是有指定默认值的，并同样指定默认值。
所有的String类型都取消Nullable属性。
去掉和时间相关的DEFAULT CURRENT_TIMESTAMP。
```

## How to code

### [GoFrame](https://github.com/gogf/gf)

> [docs](https://goframe.org/)

### [service](internal/service) Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.

Run the following command to generate code[^2]:

[^2]: [`gf gen service` docs](https://goframe.org/docs/cli/gen-service)

```bash
gf gen service
```

### Go version and dependencies updates

Go version follows the latest supported Go version of [sonic](https://github.com/bytedance/sonic).

You can easily update the dependencies by running the following command:

```bash
go get -u all
```