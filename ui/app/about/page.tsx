'use client'

import { Accordion, Breadcrumb, Card, ListGroup } from "react-bootstrap"
import Image from 'react-bootstrap/Image';
import datasource_1 from '/about/images/datasource-1.png';

export default function About() {
  return <div style={{ width: "100%" }}>
    <Breadcrumb>
      <Breadcrumb.Item href="/">Home</Breadcrumb.Item>
      <Breadcrumb.Item active>About & Help</Breadcrumb.Item>
    </Breadcrumb>

    <table>
      <tbody>
        <tr>
          <td><Image src="logo.png" style={{ width: 70 }} alt="Elasticsearch Data Collector" title="Elasticsearch Data Collector" /></td>
          <td><h2>Elasticsearch Data Collector</h2></td>
        </tr>
      </tbody>
    </table>

    <Card>
      <Card.Body>
        <Card.Title>Elasticsearch Data Collector</Card.Title>
        <Card.Text>
          You can easily transfer your data to Elasticsearch Data Collector with a few definitions and sql query.
        </Card.Text>
      </Card.Body>

      <ListGroup className="list-group-flush">
        <ListGroup.Item><b>What do you need to do this?</b></ListGroup.Item>
        <ListGroup.Item>- Create a datasource to retrieve your data</ListGroup.Item>
        <ListGroup.Item>- Write a sql query and control the result of the query with the data preview feature</ListGroup.Item>
        <ListGroup.Item>- Schedule a syncronization</ListGroup.Item>
        <ListGroup.Item>that's it. After that your data will be on the Elasticsearch Data Collector.</ListGroup.Item>
      </ListGroup>

      <Card.Body>
        <Card.Title>Starting steps</Card.Title>
      </Card.Body>

      <Accordion flush>
        <Accordion.Item eventKey="0">
          <Accordion.Header>#1 Create a datasource</Accordion.Header>
          <Accordion.Body>
            Go to the <a href="/datasources">Datasources</a> menu.
            <br />
            <br />
            <Image src="about/images/datasource-1.png" thumbnail alt="." />
            <br />
            Click "Add New" button and fill the form with your database information and save.
            <br />
            <br />
            <Image src="about/images/datasource-2.png" thumbnail alt="." />
            <br />

            <b>Postgresql Config Example</b>
            <p>
              <ul>
              <li><b>Connection String :</b>user=#USER# dbname=postgres sslmode=disable password=#PWD# host=192.168.1.60</li>
              <li><b>Driver Name :</b>postgres</li>
              </ul>
            </p>

            <b>MSSQL Config Example</b>
            <p>
              <ul>
              <li><b>Connection String :</b>server=localhost;user id=#USER#;password=#PWD#;database=master;</li>
              <li><b>Driver Name :</b>sqlserver</li>
              </ul>
            </p>

            <b>MySQL Config Example</b>
            <p>
              <ul>
              <li><b>Connection String :</b>tcp(127.0.0.1:3306)/es-data-collector?parseTime=true</li>
              <li><b>Driver Name :</b>mysql</li>
              </ul>
            </p>

            <b>Oracle Config Example</b>
            <p>
              <ul>
              <li><b>Connection String :</b>localhost:49161/xe</li>
              <li><b>Driver Name :</b>oracle</li>
              </ul>
            </p>

            You have done. Let's to the next step...
          </Accordion.Body>
        </Accordion.Item>
        <Accordion.Item eventKey="1">
          <Accordion.Header>#2 Create a Index</Accordion.Header>
          <Accordion.Body>
            Let's we create a new index.
            <br />
            <br />
            <Image src="about/images/indices-1.png" thumbnail alt="." />
            <br />
            Click "Add New" button and fill the form with your information. You should enter a suitable name lowercase and concatenated with an underscore character.
            You can enter a description for your index. You must enter a valid sql query to retrieve your data from database. After that you must select your datasource.
            At this point you can preview your sql query results with clicking preview button.
            <br />
            <br />
            <Image src="about/images/indices-2.png" thumbnail alt="." />
            <br />
            <br />
            <Image src="about/images/indices-3.png" thumbnail alt="." />
            <br />
            If everyting is OK, save the index.
          </Accordion.Body>
        </Accordion.Item>
        <Accordion.Item eventKey="2">
          <Accordion.Header>#3 Schedule data synchronization</Accordion.Header>
          <Accordion.Body>
            Go to the <a href="/indices">Indices</a> menu and click the edit button on the list.
            <br />
            <br />
            <Image src="about/images/indices-4.png" thumbnail alt="." />
            <br />
            Click the "Schedule Data Sync" link.
            <br />
            <br />
            <Image src="about/images/indices-5.png" thumbnail alt="." />
            <br />
            Enter a valid cron expression what you want your synchronization period.
            <br />
            <br />
            Enter the "Document Id Field". Document Id field must specify a unique row key in your data.
            <br />
            <br />
            You must select a "Sync Type". There are the synchronization types: "Reload All" and "Iterative". If you select the <b>reload all</b> type,
            your exist data that is on the Elasticsearch Data Collector will be deleted after synchronization. First, all the data in RDBMS will be transfered to Elasticsearch Data Collector again.
            Don't worry. Until the synchronization is completed, your exist data will be reachable. When the synchronization is completed, you can reach the new data.
            And after that the old data that is on the Elasticsearch Data Collector will be deleted.
            <br />
            <br />
            If you select the <b>Iterative type</b>, you can only retrieve the data that is changed from after last synchronization time. To do that you can use the special keyword ":#sql_last_value"
            to modify your query dynamically. For example; "select * from customers where created_at &gt;= :#sql_last_value".
            <br />
            <br />
            <Image src="about/images/indices-6.png" thumbnail alt="." />
            <br />
          </Accordion.Body>
        </Accordion.Item>

      </Accordion>

    </Card>
  </div>
}