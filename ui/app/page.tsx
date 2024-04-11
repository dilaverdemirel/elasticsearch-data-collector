'use client'
import { useEffect, useState } from 'react';
import { Card } from 'react-bootstrap';
import { CartesianGrid, Legend, Line, LineChart, Tooltip, XAxis, YAxis } from 'recharts'
import { getRecordStatistic, getStatusStatistic } from './service';

function Home() {

  const [statusStatistic, setStatusStatistic] = useState([]);
  const [recordStatistic, setRecordStatistic] = useState([]);

  const loadStatusStatistic = () => {
    getStatusStatistic()
      .then(res => {

        const result = Object.groupBy(res.data, ({ Status }) => Status);
        const result1= Object.entries(result).map(data => {
          return {
            'name': data[0],
            'data': data[1]
          }
        })

        console.log(JSON.stringify(result1))

        setStatusStatistic(result1)
      })
  }

  const loadRecordStatistic = () => {
    getRecordStatistic()
      .then(res => {
        setRecordStatistic(res.data)
      })
  }

  useEffect(() => {
    loadStatusStatistic()
    loadRecordStatistic()
  }, [])

  const colors = ["#8884d8","#82ca9d","#741d9e"]

  return (
    <div className="container">
      <div className="row">
        <div className="col-sm-12">
          <Card>
            <Card.Body>
              <Card.Title>Daily Sync Status Stats</Card.Title>

                <LineChart
                  width={900}
                  height={300}

                  margin={{
                    top: 5,
                    right: 30,
                    left: 20,
                    bottom: 5,
                  }}
                >
                  <CartesianGrid strokeDasharray="3 3" />
                  <YAxis dataKey="RecordCount"/>
                  <XAxis dataKey="Day" allowDuplicatedCategory={false}/>
                  <Tooltip />
                  <Legend />

                  {statusStatistic.map((s, index) => (
                    <Line dataKey="RecordCount" data={s.data} name={s.name} key={s.name} stroke={colors[index]}/>
                  ))}
                  
                </LineChart>

            </Card.Body>
          </Card>
        </div>

      </div>

      <div className="row" style={{ height: 10 }}></div>

      <div className="row">
        <div className="col-sm-12">
          <Card>
            <Card.Body>
              <Card.Title>Daily Sync Record Stats</Card.Title>

                <LineChart
                  width={900}
                  height={300}
                  data={recordStatistic}
                  margin={{
                    top: 5,
                    right: 30,
                    left: 20,
                    bottom: 5,
                  }}
                >
                  <CartesianGrid strokeDasharray="3 3" />
                  <XAxis dataKey="Day" />
                  <YAxis />
                  <Tooltip />
                  <Legend />
                  <Line type="monotone" dataKey="RecordCount" stroke="#8884d8" activeDot={{ r: 8 }} />
                </LineChart>

            </Card.Body>
          </Card>
        </div>

      </div>

    </div>
  )
}

export default Home