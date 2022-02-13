package artifact

import (
	"context"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/aquasecurity/fanal/analyzer"
	"github.com/aquasecurity/fanal/analyzer/config"
	"github.com/aquasecurity/fanal/artifact"
	"github.com/aquasecurity/fanal/cache"
	"github.com/aquasecurity/trivy/pkg/scanner"
)

func filesystemScanner(ctx context.Context, path string, ac cache.ArtifactCache, lac cache.LocalArtifactCache,
	_ bool, artifactOpt artifact.Option, scannerOpt config.ScannerOption) (scanner.Scanner, func(), error) {
	s, cleanup, err := initializeFilesystemScanner(ctx, path, ac, lac, artifactOpt, scannerOpt)
	if err != nil {
		return scanner.Scanner{}, func() {}, xerrors.Errorf("unable to initialize a filesystem scanner: %w", err)
	}
	return s, cleanup, nil
}

// FilesystemRun runs scan on filesystem for language-specific dependencies and config files
func FilesystemRun(ctx *cli.Context) error {
	opt, err := initOption(ctx)
	if err != nil {
		return xerrors.Errorf("option error: %w", err)
	}

	// Disable the individual package scanning
	opt.DisabledAnalyzers = analyzer.TypeIndividualPkgs

	return Run(ctx.Context, opt, filesystemScanner, initFSCache)
}

// RootfsRun runs scan on rootfs.
func RootfsRun(ctx *cli.Context) error {
	opt, err := initOption(ctx)
	if err != nil {
		return xerrors.Errorf("option error: %w", err)
	}

	// Disable the lock file scanning
	opt.DisabledAnalyzers = analyzer.TypeLockfiles

	return Run(ctx.Context, opt, filesystemScanner, initFSCache)
}
