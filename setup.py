# Copyright 2014 Richard Hawkins
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from setuptools import setup, find_packages

from transylvania import __version__ as version

name = 'transylvania'

setup(
    name=name,
    version=version,
    author='Richard Hawkins',
    author_email='hurricanerix@gmail.com',
    description='Transylvania',
    license='Apache Software License',
    keywords='opengl sdl',
    url='http://github.com/hurricanerix/transylvania',
    packages=find_packages(),
    classifiers=[
        'Development Status :: 1 - Planning',
        'License :: OSI Approved :: Apache Software License',
        'Operating System :: OS Independent',
        'Programming Language :: Python :: 2.7',
        'Environment :: Other Environment',
        ],
    install_requires=[],
    )
