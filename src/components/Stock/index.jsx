import React from "react";
import Chart from "react-google-charts";
import { Container, Row, Col, Button, ButtonGroup } from "reactstrap";
import "./stock.css";

const Stock = props => {
  return (
    <Container>
      <Row>
        <Col>
          <Chart
            height={"400px"}
            chartType="LineChart"
            loader={<div>Loading Chart</div>}
            data={props.data}
            options={{
              hAxis: { title: "Time" },
              vAxis: { title: "Stock" },
              series: { 1: { curveType: "function" } }
            }}
            rootProps={{ "data-testid": "2" }}
          />
        </Col>
      </Row>

      <Row>
        <Col sm={{ offset: 2 }}>
          <ButtonGroup>
            <Button onClick={() => props.handleSubmit(props.data[0][1], props.data[0][2], props.data[0][3], '1m')}>1M</Button>
            <Button onClick={() => props.handleSubmit(props.data[0][1], props.data[0][2], props.data[0][3], '1y')}>1A</Button>
          </ButtonGroup>
        </Col>
      </Row>
    </Container>
  );
};

export { Stock };
