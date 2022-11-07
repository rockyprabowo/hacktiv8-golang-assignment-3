import {Dialog, Transition} from '@headlessui/react'
import {Fragment, useState} from 'react'

export default function MyModal(): JSX.Element {
  const [isOpen, setIsOpen] = useState(false)

  function closeModal(): void {
    setIsOpen(false)
  }

  function openModal(): void {
    setIsOpen(true)
  }

  return (
      <>
        <div className="mt-6 flex justify-end">
          <button
              type="button"
              onClick={openModal}
              className="flex space-x-2 items-center rounded-md bg-slate-200 px-4 py-2 text-sm font-medium text-gray-900 hover:bg-opacity-50 focus:outline-none focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75"
          >
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5}
                 stroke="currentColor" className="w-6 h-6">
              <path strokeLinecap="round" strokeLinejoin="round"
                    d="M11.25 11.25l.041-.02a.75.75 0 011.063.852l-.708 2.836a.75.75 0 001.063.853l.041-.021M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-9-3.75h.008v.008H12V8.25z"/>
            </svg>
            <span>Attributions</span>
          </button>
        </div>

        <Transition appear show={isOpen} as={Fragment}>
          <Dialog as="div" className="relative z-10" onClose={closeModal}>
            <Transition.Child
                as={Fragment}
                enter="ease-out duration-300"
                enterFrom="opacity-0"
                enterTo="opacity-100"
                leave="ease-in duration-200"
                leaveFrom="opacity-100"
                leaveTo="opacity-0"
            >
              <div className="fixed inset-0 bg-black bg-opacity-25"/>
            </Transition.Child>

            <div className="fixed inset-0 overflow-y-auto">
              <div className="flex min-h-full items-center justify-center p-4 text-center">
                <Transition.Child
                    as={Fragment}
                    enter="ease-out duration-300"
                    enterFrom="opacity-0 scale-95"
                    enterTo="opacity-100 scale-100"
                    leave="ease-in duration-200"
                    leaveFrom="opacity-100 scale-100"
                    leaveTo="opacity-0 scale-95"
                >
                  <Dialog.Panel
                      className="w-full max-w-md transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all">
                    <Dialog.Title
                        as="h3"
                        className="text-2xl font-medium leading-snug text-gray-900"
                    >
                      Attributions
                    </Dialog.Title>
                    <div className="my-4">
                      <ul>
                        <li>Water energy icons created by <a href="https://www.flaticon.com/free-icons/water-energy"
                                                             title="water energy icons" className="font-medium">IconsNova
                          - Flaticon</a></li>
                        <li>Wind turbine icons created by <a href="https://www.flaticon.com/free-icons/wind-turbine"
                                                             title="wind turbine icons" className="font-medium">Freepik
                          - Flaticon</a></li>
                        <li>Water level icons created by <a href="https://www.flaticon.com/free-icons/water-level"
                                                            title="water level icons" className="font-medium">Mehwish -
                          Flaticon</a></li>
                      </ul>
                    </div>

                    <div className="mt-4">
                      <button
                          type="button"
                          className="inline-flex justify-center rounded-md border border-transparent bg-slate-100 px-4 py-2 text-sm font-medium text-gray-900 hover:bg-blue-200 focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2"
                          onClick={closeModal}
                      >
                        Close
                      </button>
                    </div>
                  </Dialog.Panel>
                </Transition.Child>
              </div>
            </div>
          </Dialog>
        </Transition>
      </>
  )
}
