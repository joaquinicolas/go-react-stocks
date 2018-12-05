import React, { Component } from "react";
import Chart from "react-google-charts";
import { Container, Row, Col, Button, ButtonGroup } from "reactstrap";
import "./stock.css";

const Stock = props => {
  console.log("ok");

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
            <Button>1M</Button>
            <Button>1A</Button>
          </ButtonGroup>
        </Col>
      </Row>
    </Container>
  );
};

export { Stock };
