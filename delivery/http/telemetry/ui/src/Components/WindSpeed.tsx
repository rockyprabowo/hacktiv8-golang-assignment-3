import {BuildTelemetryBackgroundCSS} from '../Helpers/helpers'
import windPowerIcon from '../Images/wind-power.png'

const WindSpeed = (props: WindSpeedProps): JSX.Element => {
    return (
        <div
            className={'transition duration-300 flex justify-between items-center rounded-lg p-6 ' + BuildTelemetryBackgroundCSS(props.value, props.stateMapper)}>
            <div>
                <img src={windPowerIcon} alt="Water Level Icon" className='h-24'/>
            </div>
            <div className='flex flex-col items-end gap-2'>
                <div className='font-semibold text-lg'>Wind Speed</div>
                <div><span className='text-5xl'>{props.value}</span> <span className=''>m/s</span></div>
            </div>
        </div>
    )
}

interface WindSpeedProps {
    value: number | undefined
    stateMapper: ValueStateMapperFunc
}

export default WindSpeed
