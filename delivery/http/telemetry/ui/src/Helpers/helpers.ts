function BuildTelemetryBackgroundCSS(value: number | undefined, stateMapper: ValueStateMapperFunc): string {
    if (value != null) {
        switch (stateMapper(value)) {
            case 'safe':
                return 'bg-green-200'
            case 'warning':
                return 'bg-amber-200'
            case 'danger':
                return 'bg-red-200'
            default:
                return 'bg-slate-200'
        }
    }
    return 'bg-slate-200'
}

function BuildComparatorFunc(comparator: string): ValueStateFunc {
    const comparatorParam = comparator.split(' ')
    if (comparatorParam.length < 2 || comparatorParam.length > 3) {
        throw Error('Invalid comparison parameter, expected at least two or at most three paramater')
    }
    if (comparatorParam.length === 2) {
        const comparatorLead = comparatorParam[0]
        const comparatorValue = +comparatorParam[1]
        switch (comparatorLead) {
            case '<':
                return (value: number) => value < comparatorValue
            case '<=':
                return (value: number) => value <= comparatorValue
            case '>':
                return (value: number) => value > comparatorValue
            case '>=':
                return (value: number) => value >= comparatorValue
            case '==':
                return (value: number) => value === comparatorValue
            case '!=':
                return (value: number) => value !== comparatorValue
        }
    }
    if (comparatorParam.length === 3) {
        const comparatorLead = comparatorParam[0]
        const comparatorValueLeft = +comparatorParam[1]
        const comparatorValueRight = +comparatorParam[2]
        switch (comparatorLead) {
            case '><':
                return (value: number) => value > comparatorValueLeft && value < comparatorValueRight
            case '>=<=':
                return (value: number) => value >= comparatorValueLeft && value <= comparatorValueRight
            case '<>':
                return (value: number) => value < comparatorValueLeft && value > comparatorValueRight
            case '<=>=':
                return (value: number) => value <= comparatorValueLeft && value >= comparatorValueRight
        }
    }
    throw Error(`Unparsable comparator: ${comparator}`)
}

function StateFuncBuilder(states: States): ValueStateMapperFunc {
    const stateFuncs: ValueStateFuncs = {
        safe: BuildComparatorFunc(states.safe),
        warning: BuildComparatorFunc(states.warning),
        danger: BuildComparatorFunc(states.danger)
    }

    return (value: number): string => {
        for (const key of Object.keys(states)) {
            if (stateFuncs[key](value)) {
                return key
            }
        }
        return 'default'
    }
}

export {BuildTelemetryBackgroundCSS, BuildComparatorFunc, StateFuncBuilder}