import React from "react";
import { Form, FormGroup, Label, Input, Button, Row, Col } from "reactstrap";

export class SymbolForm extends React.Component {
  constructor(props) {
    super(props);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSubmit(e) {
    e.preventDefault();
    this.props.handleSubmit(
      this.symbol1.value,
      this.symbol2.value,
      this.symbol3.value
    );
  }

  render() {
    return (
      <Row>
        <Col sm="12" md={{ size: 6, offset: 3 }} className="Form">
          <Form onSubmit={this.handleSubmit}>
            <FormGroup>
              <Label>Stocks ID-1:</Label>
              <Input
                placeholder="Enter the first symbol"
                innerRef={input => (this.symbol1 = input)}
              >
                {" "}
              </Input>
            </FormGroup>
            <FormGroup>
              <Label>Stocks ID-2:</Label>
              <Input
                placeholder="Enter the second symbol"
                innerRef={input => (this.symbol2 = input)}
              >
                {" "}
              </Input>
            </FormGroup>
            <FormGroup>
              <Label>Stocks ID-3:</Label>
              <Input
                placeholder="Enter the third symbol"
                innerRef={input => (this.symbol3 = input)}
              >
                {" "}
              </Input>
            </FormGroup>
            <FormGroup>
              <Button>Go</Button>{" "}
            </FormGroup>
          </Form>
        </Col>
      </Row>
    );
  }
}
