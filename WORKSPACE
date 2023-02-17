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

# The following are repositories based on Go dependencies for the nativeapp target automatically added by `gazelle update-repos -from_file=nativeapp/go.mod` command
go_repository(
    name = "com_eliasnaur_font",
    importpath = "eliasnaur.com/font",
    sum = "h1:djFprmHZgrSepsHAIRMp5UJn3PzsoTg9drI+BDmif5Q=",
    version = "v0.0.0-20220124212145-832bb8fc08c3",
)

go_repository(
    name = "com_github_benoitkugler_pstokenizer",
    importpath = "github.com/benoitkugler/pstokenizer",
    sum = "h1:XXpZKCZtl1kkWsI3PXEazsHPGPGa5whY7BSE09MRoRs=",
    version = "v1.0.0",
)

go_repository(
    name = "com_github_benoitkugler_textlayout",
    importpath = "github.com/benoitkugler/textlayout",
    sum = "h1:2ehWXEkgb6RUokTjXh1LzdGwG4dRP6X3dqhYYDYhUVk=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_benoitkugler_textlayout_testdata",
    importpath = "github.com/benoitkugler/textlayout-testdata",
    sum = "h1:AvFxBxpfrQd8v55qH59mZOJOQjtD6K2SFe9/HvnIbJk=",
    version = "v0.1.1",
)

go_repository(
    name = "com_github_go_gl_glfw_v3_3_glfw",
    importpath = "github.com/go-gl/glfw/v3.3/glfw",
    sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
    version = "v0.0.0-20200222043503-6f7a984d4dc4",
)

go_repository(
    name = "com_github_go_text_typesetting",
    importpath = "github.com/go-text/typesetting",
    sum = "h1:iOA0HmtpANn48hX2nlDNMu0VVaNza35HJG0WeetBVzQ=",
    version = "v0.0.0-20221214153724-0399769901d5",
)

go_repository(
    name = "com_github_google_go_cmp",
    importpath = "github.com/google/go-cmp",
    sum = "h1:e6P7q2lk1O+qJJb4BtCQXlK8vWEO8V1ZeuEdJNOqZyg=",
    version = "v0.5.8",
)

go_repository(
    name = "com_github_jezek_xgb",
    importpath = "github.com/jezek/xgb",
    sum = "h1:s2rRzAV8KQRlpsYA7Uyxoidv1nodMF0m6dIG6FhhVLQ=",
    version = "v1.0.0",
)

go_repository(
    name = "com_shuralyov_dmitri_gpu_mtl",
    importpath = "dmitri.shuralyov.com/gpu/mtl",
    sum = "h1:+PdD6GLKejR9DizMAKT5DpSAkKswvZrurk1/eEt9+pw=",
    version = "v0.0.0-20201218220906-28db891af037",
)

go_repository(
    name = "org_gioui",
    importpath = "gioui.org",
    sum = "h1:nhn8E5qDe8MVjc1bFvRHXAH7dkWDkmlbL/8dbLkCye8=",
    version = "v0.0.0-20230128030432-db6b4de0f71b",
)

go_repository(
    name = "org_gioui_cpu",
    importpath = "gioui.org/cpu",
    sum = "h1:AGDDxsJE1RpcXTAxPG2B4jrwVUJGFDjINIPi1jtO6pc=",
    version = "v0.0.0-20210817075930-8d6a761490d2",
)

go_repository(
    name = "org_gioui_shader",
    importpath = "gioui.org/shader",
    sum = "h1:cvZmU+eODFR2545X+/8XucgZdTtEjR3QWW6W65b0q5Y=",
    version = "v1.0.6",
)

go_repository(
    name = "org_golang_x_exp",
    importpath = "golang.org/x/exp",
    sum = "h1:sBdrWpxhGDdTAYNqbgBLAR+ULAPPhfgncLr1X0lyWtg=",
    version = "v0.0.0-20221012211006-4de253d81b95",
)

go_repository(
    name = "org_golang_x_exp_shiny",
    importpath = "golang.org/x/exp/shiny",
    sum = "h1:ryT6Nf0R83ZgD8WnFFdfI8wCeyqgdXWN4+CkFVNPAT0=",
    version = "v0.0.0-20220827204233-334a2380cb91",
)

go_repository(
    name = "org_golang_x_image",
    importpath = "golang.org/x/image",
    sum = "h1:/eM0PCrQI2xd471rI+snWuu251/+/jpBpZqir2mPdnU=",
    version = "v0.0.0-20220722155232-062f8c9fd539",
)

go_repository(
    name = "org_golang_x_mobile",
    importpath = "golang.org/x/mobile",
    sum = "h1:kgfVkAEEQXXQ0qc6dH7n6y37NAYmTFmz0YRwrRjgxKw=",
    version = "v0.0.0-20201217150744-e6ae53a27f4f",
)

go_repository(
    name = "org_golang_x_mod",
    importpath = "golang.org/x/mod",
    sum = "h1:6zppjxzCulZykYSLyVDYbneBfbaBIQPYMevg0bEwv2s=",
    version = "v0.6.0-dev.0.20220419223038-86c51ed26bb4",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:p9UgmWI9wKpfYmgaV/IZKGdXc5qEK45tDwwwDyjS26I=",
    version = "v0.0.0-20210510120150-4163338589ed",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:UiNENfZ8gDvpiWw7IpOMQ27spWmThO1RwwdQVbJahJM=",
    version = "v0.0.0-20220825204002-c680a09ffe64",
)

go_repository(
    name = "org_golang_x_term",
    importpath = "golang.org/x/term",
    sum = "h1:v+OssWQX+hTHEmOBgwxdZxK4zHq3yOs8F9J7mk0PY8E=",
    version = "v0.0.0-20201126162022-7de9c90e9dd1",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:olpwvP2KacW1ZWvsR7uQhoyTYvKAupfQrRGBFM352Gk=",
    version = "v0.3.7",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:VveCTK38A2rkS8ZqFY25HIDFscX5X9OoEhJd3quQmXU=",
    version = "v0.1.12",
)
