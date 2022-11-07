import {BuildTelemetryBackgroundCSS} from '../Helpers/helpers'
import seaLevelIcon from '../Images/sea-level.png'

const WaterLevel = (props: WaterLevelProps): JSX.Element => {
    return (
        <div
            className={'transition duration-300 flex justify-between items-center rounded-lg p-6 ' + BuildTelemetryBackgroundCSS(props.value, props.stateMapper)}>
            <div>
                <img src={seaLevelIcon} alt="Sea Level Icon" className='h-24'/>
            </div>
            <div className='flex flex-col items-end gap-2'>
                <div className='font-semibold text-lg'>Water Level</div>
                <div><span className='text-5xl'>{props.value}</span> <span className='text-gray-800'>m</span></div>
            </div>
        </div>
    )
}

interface WaterLevelProps {
    value: number | undefined
    stateMapper: ValueStateMapperFunc
}

export default WaterLevel
