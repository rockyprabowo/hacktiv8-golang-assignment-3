interface NetErrorProps {
  errorMessage: string
}

const NetError = (props: NetErrorProps): JSX.Element => {
  return (
    <h3 className='font-semibold'>Error: {props.errorMessage}</h3>
  )
}

export default NetError
