# Changelog

## [0.54.1](https://github.com/aquasecurity/trivy/compare/v0.54.0...v0.54.1) (2024-07-31)


### Bug Fixes

* **flag:** incorrect behavior for deprected flag `--clear-cache` [backport: release/v0.54] ([#7285](https://github.com/aquasecurity/trivy/issues/7285)) ([334a1c2](https://github.com/aquasecurity/trivy/commit/334a1c293bb3d490af2a6d80732f399efaac22f7))
* **java:** Return error when trying to find a remote pom to avoid segfault [backport: release/v0.54] ([#7283](https://github.com/aquasecurity/trivy/issues/7283)) ([f61725c](https://github.com/aquasecurity/trivy/commit/f61725c28b56d80fb46395479842a2ab0c517c5f))
* **plugin:** do not call GitHub content API for releases and tags [backport: release/v0.54] ([#7279](https://github.com/aquasecurity/trivy/issues/7279)) ([a7b7117](https://github.com/aquasecurity/trivy/commit/a7b7117fe2c9608e990b42e702cc83675c48f888))

## [0.54.0](https://github.com/aquasecurity/trivy/compare/v0.53.0...v0.54.0) (2024-07-30)


### Features

* add `log.FilePath()` function for logger ([#7080](https://github.com/aquasecurity/trivy/issues/7080)) ([1f5f348](https://github.com/aquasecurity/trivy/commit/1f5f34895823fae81bf521fc939bee743a50e304))
* add openSUSE tumbleweed detection and scanning ([#6965](https://github.com/aquasecurity/trivy/issues/6965)) ([17b5dbf](https://github.com/aquasecurity/trivy/commit/17b5dbfa12180414b87859c6c46bfe6cc5ecf7ba))
* **cli:** rename `--vuln-type` flag to `--pkg-types` flag ([#7104](https://github.com/aquasecurity/trivy/issues/7104)) ([7cbdb0a](https://github.com/aquasecurity/trivy/commit/7cbdb0a0b5dff33e506e1c1f3119951fa241b432))
* **mariner:** Add support for Azure Linux ([#7186](https://github.com/aquasecurity/trivy/issues/7186)) ([5cbc452](https://github.com/aquasecurity/trivy/commit/5cbc452a09822d1bf300ead88f0d613d4cf0349a))
* **misconf:** enabled China configuration for ACRs ([#7156](https://github.com/aquasecurity/trivy/issues/7156)) ([d1ec89d](https://github.com/aquasecurity/trivy/commit/d1ec89d1db4b039f0e31076ccd1ca969fb15628e))
* **nodejs:** add license parser to pnpm analyser ([#7036](https://github.com/aquasecurity/trivy/issues/7036)) ([03ac93d](https://github.com/aquasecurity/trivy/commit/03ac93dc208f1b40896f3fa11fa1d45293176dca))
* **sbom:** add image labels into `SPDX` and `CycloneDX` reports ([#7257](https://github.com/aquasecurity/trivy/issues/7257)) ([4a2f492](https://github.com/aquasecurity/trivy/commit/4a2f492c6e685ff577fb96a7006cd0c43755baf4))
* **sbom:** add vulnerability support for SPDX formats ([#7213](https://github.com/aquasecurity/trivy/issues/7213)) ([efb1f69](https://github.com/aquasecurity/trivy/commit/efb1f6938321eec3529ef4fea6608261f6771ae0))
* share build-in rules ([#7207](https://github.com/aquasecurity/trivy/issues/7207)) ([bff317c](https://github.com/aquasecurity/trivy/commit/bff317c77bf4a5f615a80d9875d129213bd52f6d))
* **vex:** retrieve VEX attestations from OCI registries ([#7249](https://github.com/aquasecurity/trivy/issues/7249)) ([c2fd2e0](https://github.com/aquasecurity/trivy/commit/c2fd2e0d89567a0ccd996dda8790f3c3305ea6f7))
* **vex:** VEX Repository support ([#7206](https://github.com/aquasecurity/trivy/issues/7206)) ([88ba460](https://github.com/aquasecurity/trivy/commit/88ba46047c93e6046292523ae701de774dfdc4dc))
* **vuln:** add `--pkg-relationships` ([#7237](https://github.com/aquasecurity/trivy/issues/7237)) ([5c37361](https://github.com/aquasecurity/trivy/commit/5c37361600d922db27dd594b2a80c010a19b3a6e))


### Bug Fixes

* Add dependencyManagement exclusions to the child exclusions ([#6969](https://github.com/aquasecurity/trivy/issues/6969)) ([dc68a66](https://github.com/aquasecurity/trivy/commit/dc68a662a701980d6529f61a65006f1e4728a3e5))
* add missing platform and type to spec ([#7149](https://github.com/aquasecurity/trivy/issues/7149)) ([c8a7abd](https://github.com/aquasecurity/trivy/commit/c8a7abd3b508975fcf10c254d13d1a2cd42da657))
* **cli:** error on missing config file ([#7154](https://github.com/aquasecurity/trivy/issues/7154)) ([7fa5e7d](https://github.com/aquasecurity/trivy/commit/7fa5e7d0ab67f20d434b2922725988695e32e6af))
* close file when failed to open gzip ([#7164](https://github.com/aquasecurity/trivy/issues/7164)) ([2a577a7](https://github.com/aquasecurity/trivy/commit/2a577a7bae37e5731dceaea8740683573b6b70a5))
* **dotnet:** don't include non-runtime libraries into report for `*.deps.json` files ([#7039](https://github.com/aquasecurity/trivy/issues/7039)) ([5bc662b](https://github.com/aquasecurity/trivy/commit/5bc662be9a8f072599f90abfd3b400c8ab055ed6))
* **dotnet:** show `nuget package dir not found` log only when checking `nuget` packages ([#7194](https://github.com/aquasecurity/trivy/issues/7194)) ([d76feba](https://github.com/aquasecurity/trivy/commit/d76febaee107c645e864da0f4d74a8f6ae4ad232))
* ignore nodes when listing permission is not allowed ([#7107](https://github.com/aquasecurity/trivy/issues/7107)) ([25f8143](https://github.com/aquasecurity/trivy/commit/25f8143f120965c636c5ea8386398b211b082398))
* **java:** avoid panic if deps from `pom` in `it` dir are not found ([#7245](https://github.com/aquasecurity/trivy/issues/7245)) ([4e54a7e](https://github.com/aquasecurity/trivy/commit/4e54a7e84c33c1be80c52c6db78c634bc3911715))
* **java:** use `go-mvn-version` to remove `Package` duplicates ([#7088](https://github.com/aquasecurity/trivy/issues/7088)) ([a7a304d](https://github.com/aquasecurity/trivy/commit/a7a304d53e1ce230f881c28c4f35885774cf3b9a))
* **misconf:** do not evaluate TF when a load error occurs ([#7109](https://github.com/aquasecurity/trivy/issues/7109)) ([f27c236](https://github.com/aquasecurity/trivy/commit/f27c236d6e155cb366aeef619b6ea96d20fb93da))
* **nodejs:** detect direct dependencies when using `latest` version for files `yarn.lock` + `package.json` ([#7110](https://github.com/aquasecurity/trivy/issues/7110)) ([54bb8bd](https://github.com/aquasecurity/trivy/commit/54bb8bdfb934d114b5570005853bf4bc0d40c609))
* **report:** hide empty table when all secrets/license/misconfigs are ignored ([#7171](https://github.com/aquasecurity/trivy/issues/7171)) ([c3036de](https://github.com/aquasecurity/trivy/commit/c3036de6d7719323d306a9666ccc8d928d936f9a))
* **secret:** skip regular strings contain secret patterns ([#7182](https://github.com/aquasecurity/trivy/issues/7182)) ([174b1e3](https://github.com/aquasecurity/trivy/commit/174b1e3515a6394cf8d523216d6267c1aefb820a))
* **secret:** trim excessively long lines ([#7192](https://github.com/aquasecurity/trivy/issues/7192)) ([92b13be](https://github.com/aquasecurity/trivy/commit/92b13be668bd20f8e9dac2f0cb8e5a2708b9b3b5))
* **secret:** update length of `hugging-face-access-token` ([#7216](https://github.com/aquasecurity/trivy/issues/7216)) ([8c87194](https://github.com/aquasecurity/trivy/commit/8c87194f0a6b194bc5d340c8a65bd99a3132d973))
* **server:** pass license categories to options ([#7203](https://github.com/aquasecurity/trivy/issues/7203)) ([9d52018](https://github.com/aquasecurity/trivy/commit/9d5201808da89607ae43570bdf1f335b482a6b79))


### Performance Improvements

* **debian:** use `bytes.Index` in `emptyLineSplit` to cut allocation ([#7065](https://github.com/aquasecurity/trivy/issues/7065)) ([acbec05](https://github.com/aquasecurity/trivy/commit/acbec053c985388a26d899e73b4b7f5a6d1fa210))

## [0.53.0](https://github.com/aquasecurity/trivy/compare/v0.52.0...v0.53.0) (2024-07-01)


### ⚠ BREAKING CHANGES

* **k8s:** node-collector dynamic commands support ([#6861](https://github.com/aquasecurity/trivy/issues/6861))
* add clean subcommand ([#6993](https://github.com/aquasecurity/trivy/issues/6993))
* **aws:** Remove aws subcommand ([#6995](https://github.com/aquasecurity/trivy/issues/6995))

### Features

* add clean subcommand ([#6993](https://github.com/aquasecurity/trivy/issues/6993)) ([8d0ae1f](https://github.com/aquasecurity/trivy/commit/8d0ae1f5de72d92a043dcd6b7c164d30e51b6047))
* Add local ImageID to SARIF metadata ([#6522](https://github.com/aquasecurity/trivy/issues/6522)) ([f144e91](https://github.com/aquasecurity/trivy/commit/f144e912d34234f00b5a13b7a11a0019fa978b27))
* add memory cache backend ([#7048](https://github.com/aquasecurity/trivy/issues/7048)) ([55ccd06](https://github.com/aquasecurity/trivy/commit/55ccd06df43f6ff28685f46d215ccb70f55916d2))
* **aws:** Remove aws subcommand ([#6995](https://github.com/aquasecurity/trivy/issues/6995)) ([979e118](https://github.com/aquasecurity/trivy/commit/979e118a9e0ca8943bef9143f492d7eb1fd4d863))
* **conda:** add licenses support for `environment.yml` files ([#6953](https://github.com/aquasecurity/trivy/issues/6953)) ([654217a](https://github.com/aquasecurity/trivy/commit/654217a65485ca0a07771ea61071977894eb4920))
* **dart:** use first version of constraint for dependencies using SDK version ([#6239](https://github.com/aquasecurity/trivy/issues/6239)) ([042d6b0](https://github.com/aquasecurity/trivy/commit/042d6b08c283105c258a3dda98983b345a5305c3))
* **image:** Set User-Agent header for Trivy container registry requests ([#6868](https://github.com/aquasecurity/trivy/issues/6868)) ([9b31697](https://github.com/aquasecurity/trivy/commit/9b31697274c8743d6e5a8f7a1a05daf60cd15910))
* **java:** add support for `maven-metadata.xml` files for remote snapshot repositories. ([#6950](https://github.com/aquasecurity/trivy/issues/6950)) ([1f8fca1](https://github.com/aquasecurity/trivy/commit/1f8fca1fc77b989bb4e3ba820b297464dbdd825f))
* **java:** add support for sbt projects using sbt-dependency-lock ([#6882](https://github.com/aquasecurity/trivy/issues/6882)) ([f18d035](https://github.com/aquasecurity/trivy/commit/f18d035ae13b281c96aa4ed69ca32e507d336e66))
* **k8s:** node-collector dynamic commands support ([#6861](https://github.com/aquasecurity/trivy/issues/6861)) ([8d618e4](https://github.com/aquasecurity/trivy/commit/8d618e48a2f1b60c2e4c49cdd9deb8eb45c972b0))
* **misconf:** add metadata to Cloud schema ([#6831](https://github.com/aquasecurity/trivy/issues/6831)) ([02d5404](https://github.com/aquasecurity/trivy/commit/02d540478d495416b50d7e8b187ff9f5bba41f45))
* **misconf:** add support for AWS::EC2::SecurityGroupIngress/Egress ([#6755](https://github.com/aquasecurity/trivy/issues/6755)) ([55fa610](https://github.com/aquasecurity/trivy/commit/55fa6109cd0463fd3221aae41ca7b1d8c44ad430))
* **misconf:** API Gateway V1 support for CloudFormation ([#6874](https://github.com/aquasecurity/trivy/issues/6874)) ([8491469](https://github.com/aquasecurity/trivy/commit/8491469f0b35bd9df706a433669f5b62239d4ef3))
* **misconf:** support of selectors for all providers for Rego ([#6905](https://github.com/aquasecurity/trivy/issues/6905)) ([bc3741a](https://github.com/aquasecurity/trivy/commit/bc3741ae2c68cdd00fc0aef7e51985568b2eb78a))
* **php:** add installed.json file support ([#4865](https://github.com/aquasecurity/trivy/issues/4865)) ([edc556b](https://github.com/aquasecurity/trivy/commit/edc556b85e3554c31e19b1ece189effb9ba2be12))
* **plugin:** add support for nested archives ([#6845](https://github.com/aquasecurity/trivy/issues/6845)) ([622c67b](https://github.com/aquasecurity/trivy/commit/622c67b7647f94d0a0ca3acf711d8f847cdd8d98))
* **sbom:** migrate to `CycloneDX v1.6` ([#6903](https://github.com/aquasecurity/trivy/issues/6903)) ([09e50ce](https://github.com/aquasecurity/trivy/commit/09e50ce6a82073ba62f1732d5aa0cd2701578693))


### Bug Fixes

* **c:** don't skip conan files from `file-patterns` and scan `.conan2` cache dir ([#6949](https://github.com/aquasecurity/trivy/issues/6949)) ([38b35dd](https://github.com/aquasecurity/trivy/commit/38b35dd3c804027e7a6e6a9d3c87b7ac333896c5))
* **cli:** show info message only when --scanners is available ([#7032](https://github.com/aquasecurity/trivy/issues/7032)) ([e9fc3e3](https://github.com/aquasecurity/trivy/commit/e9fc3e3397564512038ddeca2adce0efcb3f93c5))
* **cyclonedx:** trim non-URL info for `advisory.url` ([#6952](https://github.com/aquasecurity/trivy/issues/6952)) ([417212e](https://github.com/aquasecurity/trivy/commit/417212e0930aa52a27ebdc1b9370d2943ce0f8fa))
* **debian:** take installed files from the origin layer ([#6849](https://github.com/aquasecurity/trivy/issues/6849)) ([089b953](https://github.com/aquasecurity/trivy/commit/089b953462260f01c40bdf588b2568ae0ef658bc))
* **image:** parse `image.inspect.Created` field only for non-empty values ([#6948](https://github.com/aquasecurity/trivy/issues/6948)) ([0af5730](https://github.com/aquasecurity/trivy/commit/0af5730cbe56686417389c2fad643c1bdbb33999))
* **license:** return license separation using separators  `,`, `or`, etc. ([#6916](https://github.com/aquasecurity/trivy/issues/6916)) ([52f7aa5](https://github.com/aquasecurity/trivy/commit/52f7aa54b520a90a19736703f8ea63cc20fab104))
* **misconf:** fix caching of modules in subdirectories ([#6814](https://github.com/aquasecurity/trivy/issues/6814)) ([0bcfedb](https://github.com/aquasecurity/trivy/commit/0bcfedbcaa9bbe30ee5ecade5b98e9ce3cc54c9b))
* **misconf:** fix parsing of engine links and frameworks ([#6937](https://github.com/aquasecurity/trivy/issues/6937)) ([ec68c9a](https://github.com/aquasecurity/trivy/commit/ec68c9ab4580d057720179173d58734402c92af4))
* **misconf:** handle source prefix to ignore ([#6945](https://github.com/aquasecurity/trivy/issues/6945)) ([c3192f0](https://github.com/aquasecurity/trivy/commit/c3192f061d7e84eaf38df8df7c879dc00b4ca137))
* **misconf:** parsing numbers without fraction as int ([#6834](https://github.com/aquasecurity/trivy/issues/6834)) ([8141a13](https://github.com/aquasecurity/trivy/commit/8141a137ba50b553a9da877d95c7ccb491d041c6))
* **nodejs:** fix infinite loop when package link from `package-lock.json` file is broken ([#6858](https://github.com/aquasecurity/trivy/issues/6858)) ([cf5aa33](https://github.com/aquasecurity/trivy/commit/cf5aa336e660e4c98481ebf8d15dd4e54c38581e))
* **nodejs:** fix infinity loops for `pnpm` with cyclic imports ([#6857](https://github.com/aquasecurity/trivy/issues/6857)) ([7d083bc](https://github.com/aquasecurity/trivy/commit/7d083bc890eccc3bf32765c6d7e922cab2e2ef94))
* **plugin:** respect `--insecure` ([#7022](https://github.com/aquasecurity/trivy/issues/7022)) ([3d02a31](https://github.com/aquasecurity/trivy/commit/3d02a31b44924f9e2495aae087f7ca9de3314db4))
* **purl:** add missed os types ([#6955](https://github.com/aquasecurity/trivy/issues/6955)) ([2d85a00](https://github.com/aquasecurity/trivy/commit/2d85a003b22298d1101f84559f7c6b470f2b3909))
* **python:** compare pkg names from `poetry.lock` and `pyproject.toml` in lowercase ([#6852](https://github.com/aquasecurity/trivy/issues/6852)) ([faa9d92](https://github.com/aquasecurity/trivy/commit/faa9d92cfeb8d924deda2dac583b6c97099c08d9))
* **sbom:** don't overwrite `srcEpoch` when decoding SBOM files ([#6866](https://github.com/aquasecurity/trivy/issues/6866)) ([04af59c](https://github.com/aquasecurity/trivy/commit/04af59c2906bcfc7f7970b4e8f45a90f04313170))
* **sbom:** fix panic when scanning SBOM file without root component into SBOM format ([#7051](https://github.com/aquasecurity/trivy/issues/7051)) ([3d4ae8b](https://github.com/aquasecurity/trivy/commit/3d4ae8b5be94cd9b00badeece8d86c2258b2cd90))
* **sbom:** take pkg name from `purl` for maven pkgs ([#7008](https://github.com/aquasecurity/trivy/issues/7008)) ([a76e328](https://github.com/aquasecurity/trivy/commit/a76e3286c413de3dec55394fb41dd627dfee37ae))
* **sbom:** use `purl` for `bitnami` pkg names ([#6982](https://github.com/aquasecurity/trivy/issues/6982)) ([7eabb92](https://github.com/aquasecurity/trivy/commit/7eabb92ec2e617300433445718be07ac74956454))
* **sbom:** use package UIDs for uniqueness ([#7042](https://github.com/aquasecurity/trivy/issues/7042)) ([14d71ba](https://github.com/aquasecurity/trivy/commit/14d71ba63c39e51dd4179ba2d6002b46e1816e90))
* **secret:** `Asymmetric Private Key` shouldn't start with space ([#6867](https://github.com/aquasecurity/trivy/issues/6867)) ([bb26445](https://github.com/aquasecurity/trivy/commit/bb26445e3df198df77930329f532ac5ab7a67af2))
* **suse:** Add SLES 15.6 and Leap 15.6 ([#6964](https://github.com/aquasecurity/trivy/issues/6964)) ([5ee4e9d](https://github.com/aquasecurity/trivy/commit/5ee4e9d30ea814f60fd5705361cabf2e83a47a78))
* use embedded when command path not found ([#7037](https://github.com/aquasecurity/trivy/issues/7037)) ([137c916](https://github.com/aquasecurity/trivy/commit/137c9164238ffd989a0c5ed24f23a55bbf341f6e))

## [0.52.0](https://github.com/aquasecurity/trivy/compare/v0.51.1...v0.52.0) (2024-06-03)


### Features

* Add Julia language analyzer support ([#5635](https://github.com/aquasecurity/trivy/issues/5635)) ([fecafb1](https://github.com/aquasecurity/trivy/commit/fecafb1fc5bb129c7485342a0775f0dd8bedd28e))
* add support for plugin index ([#6674](https://github.com/aquasecurity/trivy/issues/6674)) ([26faf8f](https://github.com/aquasecurity/trivy/commit/26faf8f3f04b1c5f9f81c03ffc6b2008732207e2))
* **misconf:** Add support for deprecating a check ([#6664](https://github.com/aquasecurity/trivy/issues/6664)) ([88702cf](https://github.com/aquasecurity/trivy/commit/88702cfd5918b093defc5b5580f7cbf16f5f2417))
* **misconf:** add Terraform 'removed' block to schema ([#6640](https://github.com/aquasecurity/trivy/issues/6640)) ([b7a0a13](https://github.com/aquasecurity/trivy/commit/b7a0a131a03ed49c08d3b0d481bc9284934fd6e1))
* **misconf:** register builtin Rego funcs from trivy-checks ([#6616](https://github.com/aquasecurity/trivy/issues/6616)) ([7c22ee3](https://github.com/aquasecurity/trivy/commit/7c22ee3df5ee51beb90e44428a99541b3d19ab98))
* **misconf:** resolve tf module from OpenTofu compatible registry ([#6743](https://github.com/aquasecurity/trivy/issues/6743)) ([ac74520](https://github.com/aquasecurity/trivy/commit/ac7452009bf7ca0fa8ee1de8807c792eabad405a))
* **misconf:** support for VPC resources for inbound/outbound rules ([#6779](https://github.com/aquasecurity/trivy/issues/6779)) ([349caf9](https://github.com/aquasecurity/trivy/commit/349caf96bc3dd81551d488044f1adfdb947f39fb))
* **misconf:** support symlinks inside of Helm archives ([#6621](https://github.com/aquasecurity/trivy/issues/6621)) ([4eae37c](https://github.com/aquasecurity/trivy/commit/4eae37c52b035b3576361c12f70d3d9517d0a73c))
* **nodejs:** add v9 pnpm lock file support ([#6617](https://github.com/aquasecurity/trivy/issues/6617)) ([1e08648](https://github.com/aquasecurity/trivy/commit/1e0864842e32a709941d4b4e8f521602bcee684d))
* **plugin:** specify plugin version ([#6683](https://github.com/aquasecurity/trivy/issues/6683)) ([d6dc567](https://github.com/aquasecurity/trivy/commit/d6dc56732babbc9d7f788c280a768d8648aa093d))
* **python:** add license support for `requirement.txt` files ([#6782](https://github.com/aquasecurity/trivy/issues/6782)) ([29615be](https://github.com/aquasecurity/trivy/commit/29615be85e8bfeaf5a0cd51829b1898c55fa4274))
* **python:** add line number support for `requirement.txt` files ([#6729](https://github.com/aquasecurity/trivy/issues/6729)) ([2bc54ad](https://github.com/aquasecurity/trivy/commit/2bc54ad2752aba5de4380cb92c13b09c0abefd73))
* **report:** Include licenses and secrets filtered by rego to ModifiedFindings ([#6483](https://github.com/aquasecurity/trivy/issues/6483)) ([fa3cf99](https://github.com/aquasecurity/trivy/commit/fa3cf993eace4be793f85907b42365269c597b91))
* **vex:** improve relationship support in CSAF VEX ([#6735](https://github.com/aquasecurity/trivy/issues/6735)) ([a447f6b](https://github.com/aquasecurity/trivy/commit/a447f6ba94b6f8b14177dc5e4369a788e2020d90))
* **vex:** support non-root components for products in OpenVEX ([#6728](https://github.com/aquasecurity/trivy/issues/6728)) ([9515695](https://github.com/aquasecurity/trivy/commit/9515695d45e9b5c20890e27e21e3ab45bfd4ce5f))


### Bug Fixes

* clean up golangci lint configuration ([#6797](https://github.com/aquasecurity/trivy/issues/6797)) ([62de6f3](https://github.com/aquasecurity/trivy/commit/62de6f3feba6e4c56ad3922441d5b0f150c3d6b7))
* **cli:** always output fatal errors to stderr ([#6827](https://github.com/aquasecurity/trivy/issues/6827)) ([c2b9132](https://github.com/aquasecurity/trivy/commit/c2b9132a7e933a68df4cc0eb86aab23719ded1b5))
* close APKINDEX archive file ([#6672](https://github.com/aquasecurity/trivy/issues/6672)) ([5caf437](https://github.com/aquasecurity/trivy/commit/5caf4377f3a7fcb1f6e1a84c67136ae62d100be3))
* close settings.xml ([#6768](https://github.com/aquasecurity/trivy/issues/6768)) ([9c3e895](https://github.com/aquasecurity/trivy/commit/9c3e895fcb0852c00ac03ed21338768f76b5273b))
* close testfile ([#6830](https://github.com/aquasecurity/trivy/issues/6830)) ([aa0c413](https://github.com/aquasecurity/trivy/commit/aa0c413814e8915b38d2285c6a8ba5bc3f0705b4))
* **conda:** add support `pip` deps for `environment.yml` files ([#6675](https://github.com/aquasecurity/trivy/issues/6675)) ([150a773](https://github.com/aquasecurity/trivy/commit/150a77313e980cd63797a89a03afcbc97b285f38))
* **go:** add only non-empty root modules for `gobinaries` ([#6710](https://github.com/aquasecurity/trivy/issues/6710)) ([c96f2a5](https://github.com/aquasecurity/trivy/commit/c96f2a5b3de820da37e14594dd537c3b0949ae9c))
* **go:** include only `.version`|`.ver` (no prefixes) ldflags for `gobinaries` ([#6705](https://github.com/aquasecurity/trivy/issues/6705)) ([afb4f9d](https://github.com/aquasecurity/trivy/commit/afb4f9dc4730671ba004e1734fa66422c4c86dad))
* Golang version parsing from binaries w/GOEXPERIMENT ([#6696](https://github.com/aquasecurity/trivy/issues/6696)) ([696f2ae](https://github.com/aquasecurity/trivy/commit/696f2ae0ecdd4f90303f41249924a09ace70dd78))
* include packages unless it is not needed ([#6765](https://github.com/aquasecurity/trivy/issues/6765)) ([56dbe1f](https://github.com/aquasecurity/trivy/commit/56dbe1f6768fe67fbc1153b74fde0f83eaa1b281))
* **misconf:** don't shift ignore rule related to code ([#6708](https://github.com/aquasecurity/trivy/issues/6708)) ([39a746c](https://github.com/aquasecurity/trivy/commit/39a746c77837f873e87b81be40676818030f44c5))
* **misconf:** skip Rego errors with a nil location ([#6638](https://github.com/aquasecurity/trivy/issues/6638)) ([a2c522d](https://github.com/aquasecurity/trivy/commit/a2c522ddb229f049999c4ce74ef75a0e0f9fdc62))
* **misconf:** skip Rego errors with a nil location ([#6666](https://github.com/aquasecurity/trivy/issues/6666)) ([a126e10](https://github.com/aquasecurity/trivy/commit/a126e1075a44ef0e40c0dc1e214d1c5955f80242))
* node-collector high and critical cves ([#6707](https://github.com/aquasecurity/trivy/issues/6707)) ([ff32deb](https://github.com/aquasecurity/trivy/commit/ff32deb7bf9163c06963f557228260b3b8c161ed))
* **plugin:** initialize logger ([#6836](https://github.com/aquasecurity/trivy/issues/6836)) ([728e77a](https://github.com/aquasecurity/trivy/commit/728e77a7261dc3fcda1e61e79be066c789bbba0c))
* **python:** add package name and version validation for `requirements.txt` files. ([#6804](https://github.com/aquasecurity/trivy/issues/6804)) ([ea3a124](https://github.com/aquasecurity/trivy/commit/ea3a124fc7162c30c7f1a59bdb28db0b3c8bb86d))
* **report:** hide empty tables if all vulns has been filtered ([#6352](https://github.com/aquasecurity/trivy/issues/6352)) ([3d388d8](https://github.com/aquasecurity/trivy/commit/3d388d8552ef42d4d54176309a38c1879008527b))
* **sbom:** fix panic for `convert` mode when scanning json file derived from sbom file ([#6808](https://github.com/aquasecurity/trivy/issues/6808)) ([f92ea09](https://github.com/aquasecurity/trivy/commit/f92ea096856c7c262b05bd4d31c62689ebafac82))
* use of specified context to obtain cluster name ([#6645](https://github.com/aquasecurity/trivy/issues/6645)) ([39ebed4](https://github.com/aquasecurity/trivy/commit/39ebed45f8c218509d264bd3f3ca548fc33d2b3a))


### Performance Improvements

* **misconf:** parse rego input once ([#6615](https://github.com/aquasecurity/trivy/issues/6615)) ([67c6b1d](https://github.com/aquasecurity/trivy/commit/67c6b1d473999003d682bdb42657bbf3a4a69a9c))
