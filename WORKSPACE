workspace(name = "com_github_antonhornquist_monorepo1")

# The load statement is used to import a symbol from an extension. Bazel extensions are files ending in .bzl. The first argument of load is a label identifying a .bzl file. See https://bazel.build/concepts/build-files#load
# The first part of the label is the repository name, ie. @bazel_tools//. In the typical case that a label refers to the same repository from which it is used, the repository identifier may be abbreviated as //. See https://bazel.build/concepts/labels

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")  # Loads the file tools/build_defs/repo:http.bzl from the remote repository @bazel_tools and adds the http_archive symbol to the environment. The http_archive statement downloads a Bazel repository as a compressed archive file, decompresses it, and makes its targets available for binding. See https://bazel.build/rules/lib/repo/http#http_archive

# The following snippet tells Bazel to fetch the io_bazel_rules_go repository and its dependencies. Bazel will download a recent supported Go toolchain and register it for use. See https://github.com/bazelbuild/rules_go
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "dd926a88a564a9246713a9c00b35315f54cbd46b31a26d5d8fb264c07045f05d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.38.1/rules_go-v0.38.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.38.1/rules_go-v0.38.1.zip",
    ],
)
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
go_rules_dependencies()
go_register_toolchains(version = "1.19.5")

# The following snippet tells Bazel to fetch the bazel_gazelle repository and its dependencies. See https://github.com/bazelbuild/bazel-gazelle
http_archive(
    name = "bazel_gazelle",
    sha256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098274d14a1d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
    ],
)
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
gazelle_dependencies()
