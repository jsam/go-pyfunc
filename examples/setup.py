from setuptools import setup


setup(
    name="gopyfunc",
    version="0.0.1",
    description="example python worker",
    author="Sam Dash",
    author_email="contact@justsam.io",
    url="https://github.com/jsam/go-pyfunc/",
    download_url="https://github.com/jsam/go-pyfunc/archive/main.zip",
    classifiers=[
        "Development Status :: 2 - Pre-Alpha",
        "Operating System :: POSIX :: Linux",
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3",
        "Topic :: Software Development",
        "Topic :: Utilities",
    ],
    py_modules=["worker"],
)
