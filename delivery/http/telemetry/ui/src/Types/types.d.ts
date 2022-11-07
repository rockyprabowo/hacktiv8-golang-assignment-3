interface States {
    safe: string
    warning: string
    danger: string

    [key: string]: string
}

type KeyNameFunc<T, VM> = {
    [K in keyof T]: VM
}
type ValueStateFunc = (value: number) => boolean

type ValueStateFuncs = KeyNameFunc<States, ValueStateFunc>

type ValueStateMapperFunc = (value: number) => string

interface TelemetryData {
    water: number
    wind: number
}

interface Telemetry {
    status: TelemetryData
}

interface Config {
    interval: string
    interval_ms: number
    water_states: States
    wind_states: States
}
