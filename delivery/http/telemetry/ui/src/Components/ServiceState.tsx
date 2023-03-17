import { useQuery } from '@tanstack/react-query'
import { fetchConfig } from '../Fetcher'

const intervalBaseStyle = 'ml-8 inline-flex cursor-default items-center justify-center whitespace-nowrap rounded-md border border-transparent px-4 py-2 text-base font-medium text-white shadow-sm'
const waitingStyle = 'bg-indigo-600 hover:bg-indigo-700'
const okStyle = 'bg-green-600 hover:bg-green-700'
const errorStyle = 'bg-red-600 hover:bg-red-700'

const ServiceState = (): JSX.Element => {
  const { data, isError, isLoading } = useQuery(
    ['config'],
    fetchConfig,
    {
      refetchInterval: 3000
    }
  )

  if (isLoading) {
    return (
      <div className={`${intervalBaseStyle} ${waitingStyle}`}>
          Waiting...
      </div>
    )
  }

  if (isError) {
    return (
      <div className={`${intervalBaseStyle} ${errorStyle}`}>
          Error!
      </div>
    )
  }

  return (
    <div className={`${intervalBaseStyle} ${okStyle}`}>
        Interval: {data?.interval}
    </div>
  )
}

export default ServiceState
