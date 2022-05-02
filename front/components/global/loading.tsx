import { FC } from 'react'
import { Spinner } from 'react-bootstrap'
import styles from './loading.module.css'

type Props = {
  show: boolean
}

const Loading: FC<Props> = props => {
  return (
    <>
    {props.show ?
    <div className={styles.container}>
      <Spinner animation="grow" />
    </div> : <></>}
    </>
  )
}

export default Loading