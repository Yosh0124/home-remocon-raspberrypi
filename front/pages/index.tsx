import type { NextPage } from 'next'
import Head from 'next/head'
import { useState } from 'react'
import { Card, Row, Col, Button } from 'react-bootstrap'
import Loading from '../components/global/loading'
import styles from '../styles/Home.module.css'
import { ControlLight } from '../lib/api/light'
import { ControlAir } from '../lib/api/air'

const Home: NextPage = () => {

  const [showLoading, setShowLoading] = useState<boolean>(false)

  const onChangeLight = (action: string) => {
    setShowLoading(true)
    ControlLight(action)
    .then(res => {
      if (res.ok) {
        return res.json()
      } else {
        throw new Error(res.status + ' ' + res.statusText)
      }
    })
    .then((data) => {
      setShowLoading(false)
    })
    .catch((err) => {
      setShowLoading(false)
      console.error(err)
    });
  }
  
  const onChangeAir = (action: string) => {
    setShowLoading(true)
    ControlAir(action)
    .then(res => {
      if (res.ok) {
        return res.json()
      } else {
        throw new Error(res.status + ' ' + res.statusText)
      }
    })
    .then((data) => {
      setShowLoading(false)
    })
    .catch((err) => {
      setShowLoading(false)
      console.error(err)
    });
  }

  return (
    <>
    <div className={styles.container}>
      <Head>
        <title>リモコン</title>
        <meta name="description" content="Remocon App" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Card className={'mb-3'}>
        <Card.Body>
          <Card.Title>
            調光
          </Card.Title>
          <Row>
            <Col><Button variant="info" size="lg" className={styles.button} onClick={()=>{onChangeLight('on')}}>ON</Button></Col>
            <Col><Button variant="warning" size="lg" className={styles.button} onClick={()=>{onChangeLight('small')}}>NIGHT</Button></Col>
            <Col><Button variant="secondary" size="lg" className={styles.button} onClick={()=>{onChangeLight('off')}}>OFF</Button></Col>
          </Row>
        </Card.Body>
      </Card>
      <Card className={'mb-3'}>
        <Card.Body>
          <Card.Title>
            エアコン
          </Card.Title>
          <Row>
            <Col><Button variant="info" size="lg" className={styles.button} onClick={()=>{onChangeAir('cooler')}}>COOLER</Button></Col>
            <Col><Button variant="danger" size="lg" className={styles.button} onClick={()=>{onChangeAir('heater')}}>HEATER</Button></Col>
            <Col><Button variant="secondary" size="lg" className={styles.button} onClick={()=>{onChangeAir('off')}}>OFF</Button></Col>
          </Row>
        </Card.Body>
      </Card>
    </div>
    <Loading show={showLoading} />
    </>
  )
}

export default Home
