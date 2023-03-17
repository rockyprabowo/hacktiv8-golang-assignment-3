import ServiceState from './ServiceState'
import waterEnergyIcon from '../Images/water-energy.png'

const Navigation = (): JSX.Element => {
  return (
    <div className='relative bg-slate-50 border-b-2 border-gray-100'>
        <div className='mx-auto max-w-8xl px-4 sm:px-6'>
            <div className='flex items-center justify-between py-6 md:justify-start md:space-x-10'>
                <div className='flex justify-start items-center lg:w-0 lg:flex-1'>
                    <img src={waterEnergyIcon} alt="Icon" className='h-16 p-2'/>
                    <span className='font-semibold text-lg ml-4'>Telemetry Dashboard</span>
                </div>
                <div className='items-center justify-end md:flex md:flex-1 lg:w-0'>
                    <ServiceState/>
                </div>
            </div>
        </div>
    </div>
  )
}

export default Navigation
