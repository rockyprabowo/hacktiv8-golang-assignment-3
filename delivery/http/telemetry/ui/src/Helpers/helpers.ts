function BuildTelemetryBackgroundCSS (value: number | undefined, stateMapper: ValueStateMapperFunc): string {
  if (value != null) {
    return `bg-${stateMapper(value)}-200`
  }
  return 'bg-slate-200'
}

function BuildComparatorFunc (comparator: string): ValueStateFunc {
  const comparatorParam = comparator.split(' ')
  if (comparatorParam.length < 2 || comparatorParam.length > 3) {
    throw new Error('Invalid comparison parameter, expected at least two or at most three parameter')
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
  throw new Error(`Couldn't parse comparator: ${comparator}`)
}

function StateFuncBuilder (states: States): ValueStateMapperFunc {
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

export { BuildTelemetryBackgroundCSS, BuildComparatorFunc, StateFuncBuilder }
