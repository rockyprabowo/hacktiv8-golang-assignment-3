import { Attributions, Navigation, Telemetry } from '../Components'
import {
  QueryClient,
  QueryClientProvider
} from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'

const App = (): JSX.Element => {
  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
        <Navigation/>
        <div className="m-6 max-w-4xl lg:mx-auto h-max">
            <Telemetry/>
            <Attributions/>
        </div>
        <ReactQueryDevtools initialIsOpen={false}/>
    </QueryClientProvider>
  )
}

export default App
