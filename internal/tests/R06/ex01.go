package R06

import (
	"path/filepath"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/logger"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var Tests01 = `
#[cfg(test)]
mod shortinette_tests_rust_0601 {
    use super::*;

    #[test]
    fn transmute_both() {
        let iron = 0x01234567;
        assert_eq!(PhilosopherStone.transmute_iron(iron), [0x4567, 0x0123]);
        let mercure = 0x0123456789ABCDEF;
        assert_eq!(
            PhilosopherStone.transmute_mercure(mercure),
            [0xCDEF, 0x89AB, 0x4567, 0x0123],
        );
    }

    #[test]
    fn transmute_metal() {
        let nugget: GoldNugget = 0x1234;
        assert_eq!(PhilosopherStone.transmute_metal(&nugget), &[0x1234]);

        let iron: Iron = 0x12345678;
        assert_eq!(PhilosopherStone.transmute_metal(&iron), &[0x5678, 0x1234]);
        let mercure: Mercure = 0x0123456789ABCDEF;
        assert_eq!(
            PhilosopherStone.transmute_metal(&mercure),
            &[0xCDEF, 0x89AB, 0x4567, 0x0123],
        );
    }

    #[test]
    fn transmute_iron_edge_values() {
        let iron_min: Iron = 0x00000000;
        assert_eq!(PhilosopherStone.transmute_iron(iron_min), [0x0000, 0x0000]);

        let iron_max: Iron = 0xFFFFFFFF;
        assert_eq!(PhilosopherStone.transmute_iron(iron_max), [0xFFFF, 0xFFFF]);
    }

    #[test]
    fn transmute_mercure_edge_values() {
        let mercure_min: Mercure = 0x0000000000000000;
        assert_eq!(
            PhilosopherStone.transmute_mercure(mercure_min),
            [0x0000, 0x0000, 0x0000, 0x0000]
        );

        let mercure_max: Mercure = 0xFFFFFFFFFFFFFFFF;
        assert_eq!(
            PhilosopherStone.transmute_mercure(mercure_max),
            [0xFFFF, 0xFFFF, 0xFFFF, 0xFFFF]
        );
    }

    #[test]
    fn transmute_iron_random_values() {
        let iron: Iron = 0xAABBCCDD;
        assert_eq!(PhilosopherStone.transmute_iron(iron), [0xCCDD, 0xAABB]);
    }

    #[test]
    fn transmute_mercure_random_values() {
        let mercure: Mercure = 0x1122334455667788;
        assert_eq!(
            PhilosopherStone.transmute_mercure(mercure),
            [0x7788, 0x5566, 0x3344, 0x1122]
        );
    }

    #[test]
    fn transmute_metal_goldnugget() {
        let nugget: GoldNugget = 0xABCD;
        assert_eq!(PhilosopherStone.transmute_metal(&nugget), &[0xABCD]);
    }

    #[test]
    fn transmute_metal_iron() {
        let iron: Iron = 0x89ABCDEF;
        assert_eq!(PhilosopherStone.transmute_metal(&iron), &[0xCDEF, 0x89AB]);
    }

    #[test]
    fn transmute_metal_mercure() {
        let mercure: Mercure = 0x1234567890ABCDEF;
        assert_eq!(
            PhilosopherStone.transmute_metal(&mercure),
            &[0xCDEF, 0x90AB, 0x5678, 0x1234]
        );
    }

    #[test]
    fn transmute_iron_zero() {
        let iron: Iron = 0x00000000;
        assert_eq!(PhilosopherStone.transmute_iron(iron), [0x0000, 0x0000]);
    }

    #[test]
    fn transmute_mercure_zero() {
        let mercure: Mercure = 0x0000000000000000;
        assert_eq!(
            PhilosopherStone.transmute_mercure(mercure),
            [0x0000, 0x0000, 0x0000, 0x0000]
        );
    }

    #[test]
    fn transmute_metal_zero_goldnugget() {
        let nugget: GoldNugget = 0x0000;
        assert_eq!(PhilosopherStone.transmute_metal(&nugget), &[0x0000]);
    }

    #[test]
    fn transmute_metal_zero_iron() {
        let iron: Iron = 0x00000000;
        assert_eq!(PhilosopherStone.transmute_metal(&iron), &[0x0000, 0x0000]);
    }

    #[test]
    fn transmute_metal_zero_mercure() {
        let mercure: Mercure = 0x0000000000000000;
        assert_eq!(
            PhilosopherStone.transmute_metal(&mercure),
            &[0x0000, 0x0000, 0x0000, 0x0000]
        );
    }
}
`

func ex01Test(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.RepoDirectory, exercise.TurnInDirectory)
	if err := testutils.AppendStringToFile(Tests01, exercise.TurnInFiles[0]); err != nil {
		logger.Exercise.Printf("could not write to %s: %v", exercise.TurnInFiles[0], err)
		return Exercise.InternalError(err.Error())
	}
	_, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"test"})
	if err != nil {
		return Exercise.InternalError(err.Error())
	}
	return Exercise.Passed("OK")
}

func ex01() Exercise.Exercise {
	return Exercise.NewExercise("01", "studentcode", "ex01", []string{"src/lib.rs", "Cargo.toml"}, []string{"std::mem::transmute", "std::slice::from_raw_parts"}, map[string]int{"unsafe": 101}, 10, ex01Test)
}