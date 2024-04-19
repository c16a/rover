load("@rules_jvm_external//:defs.bzl", "artifact")
load("@rules_java//java:defs.bzl", "java_test")

JUNIT5_RUNTIME_DEPS = [
    artifact("org.junit.platform:junit-platform-launcher"),
    artifact("org.junit.platform:junit-platform-reporting"),
    artifact("org.junit.platform:junit-platform-console"),
    artifact("org.junit.jupiter:junit-jupiter-engine"),
]

JUNIT5_COMPILE_DEPS = [
    artifact("org.junit.jupiter:junit-jupiter-api"),
    artifact("org.junit.jupiter:junit-jupiter-params"),
]

MOCKITO_DEPS = [
    artifact("org.mockito:mockito-core"),
]

def java_junit5_test(name, srcs, test_package, deps = [], runtime_deps = [], **kwargs):
    FILTER_KWARGS = [
        "main_class",
        "use_testrunner",
        "args",
    ]

    for arg in FILTER_KWARGS:
        if arg in kwargs.keys():
            kwargs.pop(arg)

    junit_console_args = []
    if test_package:
        junit_console_args += ["--select-package", test_package]
    else:
        fail("must specify 'test_package'")

    java_test(
        name = name,
        srcs = srcs,
        use_testrunner = False,
        main_class = "org.junit.platform.console.ConsoleLauncher",
        args = junit_console_args,
        deps = deps + JUNIT5_COMPILE_DEPS + MOCKITO_DEPS,
        runtime_deps = runtime_deps + JUNIT5_RUNTIME_DEPS,
        **kwargs
    )
