workspace(name = "gitlab_k8s_agent")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

# When updating rules_go make sure to update org_golang_x_tools dependency below by copying it from
# https://github.com/bazelbuild/rules_go/blob/master/go/private/repositories.bzl
# Also update to the same version/commit in go.mod.
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "278b7ff5a826f3dc10f04feaf0b70d48b68748ccd512d7f98bf442077f043fe3",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.41.0/rules_go-v0.41.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "29218f8e0cebe583643cbf93cae6f971be8a2484cdcfa1e45057658df8d54002",
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.32.0/bazel-gazelle-v0.32.0.tar.gz",
    ],
)

http_archive(
    name = "rules_oci",
    sha256 = "db57efd706f01eb3ce771468366baa1614b5b25f4cce99757e2b8d942155b8ec",
    strip_prefix = "rules_oci-1.0.0",
    url = "https://github.com/bazel-contrib/rules_oci/releases/download/v1.0.0/rules_oci-v1.0.0.tar.gz",
)

http_archive(
    name = "rules_pkg",
    sha256 = "8f9ee2dc10c1ae514ee599a8b42ed99fa262b757058f65ad3c384289ff70c4b8",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.9.1/rules_pkg-0.9.1.tar.gz",
        "https://github.com/bazelbuild/rules_pkg/releases/download/0.9.1/rules_pkg-0.9.1.tar.gz",
    ],
)

http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "977a0bd4593c8d4c8f45e056d181c35e48aa01ad4f8090bdb84f78dca42f47dc",
    strip_prefix = "buildtools-6.1.2",
    urls = ["https://github.com/bazelbuild/buildtools/archive/refs/tags/v6.1.2.tar.gz"],
)

http_archive(
    name = "com_github_ash2k_bazel_tools",
    sha256 = "a911dab6711bc12a00f02cc94b66ced7dc57650e382ebd4f17c9cdb8ec2cbd56",
    strip_prefix = "bazel-tools-2add5bb84c2837a82a44b57e83c7414247aed43a",
    urls = ["https://github.com/ash2k/bazel-tools/archive/2add5bb84c2837a82a44b57e83c7414247aed43a.tar.gz"],
)

http_archive(
    name = "rules_proto",
    sha256 = "dc3fb206a2cb3441b485eb1e423165b231235a1ea9b031b4433cf7bc1fa460dd",
    strip_prefix = "rules_proto-5.3.0-21.7",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/5.3.0-21.7.tar.gz",
    ],
)

http_archive(
    name = "rules_proto_grpc",
    sha256 = "928e4205f701b7798ce32f3d2171c1918b363e9a600390a25c876f075f1efc0a",
    strip_prefix = "rules_proto_grpc-4.4.0",
    urls = ["https://github.com/rules-proto-grpc/rules_proto_grpc/releases/download/4.4.0/rules_proto_grpc-4.4.0.tar.gz"],
)

http_archive(
    name = "bazel_skylib",
    sha256 = "b8a1527901774180afc798aeb28c4634bdccf19c4d98e7bdd1ce79d1fe9aaad7",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.4.1/bazel-skylib-1.4.1.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.4.1/bazel-skylib-1.4.1.tar.gz",
    ],
)

http_archive(
    name = "bazelruby_rules_ruby",
    sha256 = "5035393cb5043d49ca9de78acb9e8c8622a193f6463a57ad02383a622b6dc663",
    strip_prefix = "rules_ruby-0.6.0",
    urls = [
        "https://github.com/bazelruby/rules_ruby/archive/v0.6.0.tar.gz",
    ],
)

http_archive(
    name = "aspect_bazel_lib",
    sha256 = "e3151d87910f69cf1fc88755392d7c878034a69d6499b287bcfc00b1cf9bb415",
    strip_prefix = "bazel-lib-1.32.1",
    url = "https://github.com/aspect-build/bazel-lib/releases/download/v1.32.1/bazel-lib-v1.32.1.tar.gz",
)

# Required for proto files. Needs to be compatible with generated code in org_golang_google_genproto.
# See https://github.com/googleapis/googleapis and https://github.com/googleapis/go-genproto.
# See https://github.com/bazelbuild/bazel-gazelle/releases/tag/v0.32.0
http_archive(
    name = "go_googleapis",
    sha256 = "b54f92060bf2a65c2671491d8c1cb8ad6aa466a29b58f08a907eb170e99cc7ec",
    strip_prefix = "googleapis-64d54ff78cfe7cdc2f4b95717dc9afa3ef32a29a",
    urls = [
        "https://github.com/googleapis/googleapis/archive/64d54ff78cfe7cdc2f4b95717dc9afa3ef32a29a.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

# See https://github.com/open-telemetry/opentelemetry-go-contrib/issues/872
go_repository(
    name = "io_opentelemetry_go_otel",
    build_directives = [
        "gazelle:go_visibility @io_opentelemetry_go_otel_metric//:__subpackages__",
    ],  # keep
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/otel",
    sum = "h1:Z7GVAX/UkAXPKsy94IU+i6thsQS4nb7LviLpnaNeW8s=",
    version = "v1.16.0",
)

# See https://github.com/open-telemetry/opentelemetry-go-contrib/issues/872
go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace",
    build_directives = [
        "gazelle:resolve go go.opentelemetry.io/otel/exporters/otlp/internal @io_opentelemetry_go_otel//exporters/otlp/internal",
        "gazelle:resolve go go.opentelemetry.io/otel/exporters/otlp/internal/envconfig @io_opentelemetry_go_otel//exporters/otlp/internal/envconfig",
    ],
    build_file_proto_mode = "disable",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace",
    sum = "h1:cbsD4cUcviQGXdw8+bo5x2wazq10SKz8hEbtCRPcU78=",
    version = "v1.16.0",
)

# See https://github.com/open-telemetry/opentelemetry-go-contrib/issues/872
go_repository(
    name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracehttp",
    build_directives = [
        "gazelle:resolve go go.opentelemetry.io/otel/exporters/otlp/internal @io_opentelemetry_go_otel//exporters/otlp/internal",
    ],
    build_file_proto_mode = "disable_global",
    importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp",
    sum = "h1:iqjq9LAB8aK++sKVcELezzn655JnBNdsDhghU4G/So8=",
    version = "v1.16.0",
)

# Copied from rules_go to keep patches in place
# https://github.com/bazelbuild/rules_go/blob/master/go/private/repositories.bzl
http_archive(
    name = "org_golang_x_tools",
    patch_args = ["-p1"],
    patches = [
        # deletegopls removes the gopls subdirectory. It contains a nested
        # module with additional dependencies. It's not needed by rules_go.
        # releaser:patch-cmd rm -rf gopls
        "@io_bazel_rules_go//third_party:org_golang_x_tools-deletegopls.patch",
        # releaser:patch-cmd gazelle -repo_root . -go_prefix golang.org/x/tools -go_naming_convention import_alias
        "@io_bazel_rules_go//third_party:org_golang_x_tools-gazelle.patch",
    ],
    sha256 = "9f20a20f29f4008d797a8be882ef82b69cf8f7f2b96dbdfe3814c57d8280fa4b",
    strip_prefix = "tools-0.7.0",
    # v0.7.0, latest as of 2023-03-27
    urls = [
        "https://github.com/golang/tools/archive/refs/tags/v0.7.0.zip",
        "https://mirror.bazel.build/github.com/golang/tools/archive/refs/tags/v0.7.0.zip",
    ],
)

load("//build:repositories.bzl", "go_repositories")

# gazelle:repository_macro build/repositories.bzl%go_repositories
go_repositories()

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

go_rules_dependencies()

go_register_toolchains(
    version = "1.20.6",
)

gazelle_dependencies()

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")
load("@com_github_ash2k_bazel_tools//buildozer:deps.bzl", "buildozer_dependencies")
load("@com_github_ash2k_bazel_tools//multirun:deps.bzl", "multirun_dependencies")
load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")
load("@rules_proto_grpc//:repositories.bzl", "rules_proto_grpc_repos", "rules_proto_grpc_toolchains")
load("@rules_proto_grpc//go:repositories.bzl", rules_proto_grpc_go_repos = "go_repos")
load("@rules_proto_grpc//doc:repositories.bzl", rules_proto_grpc_doc_repos = "doc_repos")
load("@aspect_bazel_lib//lib:repositories.bzl", "aspect_bazel_lib_dependencies")
load("@rules_oci//oci:dependencies.bzl", "rules_oci_dependencies")
load("@rules_oci//oci:repositories.bzl", "LATEST_CRANE_VERSION", "oci_register_toolchains")
load("@rules_oci//oci:pull.bzl", "oci_pull")
load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")
load("@go_googleapis//:repository_rules.bzl", "switched_rules_by_language")

rules_proto_dependencies()

rules_proto_toolchains()

rules_proto_grpc_toolchains()

rules_proto_grpc_repos()

rules_proto_grpc_go_repos()

rules_proto_grpc_doc_repos()

rules_oci_dependencies()

oci_register_toolchains(
    name = "oci",
    crane_version = LATEST_CRANE_VERSION,
)

buildifier_dependencies()

buildozer_dependencies()

multirun_dependencies()

load("@com_github_grpc_grpc//bazel:grpc_deps.bzl", "grpc_deps")

grpc_deps()

aspect_bazel_lib_dependencies()

rules_pkg_dependencies()

switched_rules_by_language(
    name = "com_google_googleapis_imports",
    go = True,
)

# Images are managed by https://gitlab.com/gitlab-org/frontend/renovate-gitlab-bot/-/tree/main/renovate/projects/gitlab-agent.config.js
# DO NOT EDIT ================ START

# nonroot from https://console.cloud.google.com/gcr/images/distroless/GLOBAL/static-debian11
oci_pull(
    name = "distroless_static_nonroot",
    digest = "sha256:9ecc53c269509f63c69a266168e4a687c7eb8c0cfd753bd8bfcaa4f58a90876f",
    image = "gcr.io/distroless/static-debian11",
    platforms = [
        "linux/amd64",
        "linux/arm",
        "linux/arm64",
    ],
)

# debug-nonroot from https://console.cloud.google.com/gcr/images/distroless/GLOBAL/base-debian11
oci_pull(
    name = "distroless_base_debug_nonroot",
    digest = "sha256:de8fb012fc630b7cdea6861442a0185213b574c71e246ddc97e9eb1d047048e7",
    image = "gcr.io/distroless/base-debian11",
    platforms = [
        "linux/amd64",
        "linux/arm",
        "linux/arm64",
    ],
)

# DO NOT EDIT ================ END
