load("@rules_go//go:def.bzl", "go_binary", "go_test")

# This was the best piece of code I have written.
# Then I added a for-loop in it.
# I know, I'm not proud of it.
def go_matrix_binary(name, distros, pure = "on", static = "on", **kwargs):
    for distro in distros:
        os = distro.split("/")[0]
        arch = distro.split("/")[1]
        go_binary(
            name = "{name}_{os}_{arch}".format(name = name, os = os, arch = arch),
            gc_linkopts = [
                "-s",
                "-w",
            ],
            goarch = arch,
            goos = os,
            pure = pure,
            static = static,
            **kwargs
        )

DISTROS = [
    "linux/amd64",  # Every other computer in the world
    "linux/arm",  # Raspberry Pi and other embedded ARM devices
    "linux/arm64",  # Amazon Graviton, Ampere Altra
    "linux/s390x",  # IBM LinuxONE
    "linux/ppc64",  # I don't know why this exists
    "linux/riscv64",  # I really hope this takes off in my life time
    "windows/amd64",  # Every other computer used by a person who hates Linux
    "windows/arm64",  # That mad lad who dared to run Windows on ARM
    "darwin/arm64",  # The Apple fanboy using his Mac Mini as a server
    "freebsd/amd64",
    "freebsd/arm64",
    "openbsd/amd64",
    "openbsd/arm64",
    "netbsd/amd64",
    "netbsd/arm64",
]
