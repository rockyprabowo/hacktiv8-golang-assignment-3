import Loading from './Loading'
import TelemetryData from './TelemetryData'
import NetError from './NetError'
import { useQuery } from '@tanstack/react-query'
import { fetchConfig, fetchData } from '../Fetcher'

const Telemetry = (): JSX.Element => {
  const { data: config } = useQuery(
    ['config'],
    fetchConfig
  )

  const interval = config?.interval_ms

  const { data, isError, isLoading } = useQuery(
    ['data'],
    fetchData,
    {
      enabled: interval != null,
      refetchInterval: interval
    }
  )

  if (isError) {
    return (
      <div className='p-8 rounded-lg bg-red-500 text-gray-50'>
          <NetError errorMessage='Service unavailable'/>
      </div>
    )
  }

  if (isLoading) {
    return (
      <div className='p-8 rounded-lg grid grid-rows-1 bg-slate-200'>
          <Loading/>
      </div>
    )
  }

  return (
    <TelemetryData state={data} config={config}/>
  )
}

export default Telemetry
