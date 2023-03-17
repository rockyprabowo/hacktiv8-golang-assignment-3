import { StateFuncBuilder } from '../Helpers/helpers'
import WaterLevel from './WaterLevel'
import WindSpeed from './WindSpeed'

interface TelemetryDataProps {
  state?: Telemetry
  config?: Config
}

const TelemetryData = (props: TelemetryDataProps): JSX.Element => {
  if (props.config != null) {
    const getWaterState = StateFuncBuilder(props.config.water_states)
    const getWindState = StateFuncBuilder(props.config.wind_states)

    return (
      <div className='grid grid-rows-1 md:grid-cols-2 gap-6'>
          <WaterLevel value={props.state?.status.water} stateMapper={getWaterState}/>
          <WindSpeed value={props.state?.status.wind} stateMapper={getWindState}/>
      </div>
    )
  }
  return (<></>)
}

export default TelemetryData
