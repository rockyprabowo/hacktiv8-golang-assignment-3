import axios from 'axios'

const fetchData = async (): Promise<Telemetry> => {
  const { data } = await axios.get<Telemetry>('/api/v1/telemetry')
  return data
}

const fetchConfig = async (): Promise<Config> => {
  const { data } = await axios.get<Config>('/api/v1/telemetry/config')
  return data
}

export { fetchData, fetchConfig }
