import React from 'react';
import { fireEvent, Event } from '../util';
import { Footer } from './Footer';
import { Form } from './Form';
import { TimeblockType } from './Timeblock';
import { TimeblockContainer } from './TimeblockContainer';

type Props = {
  message: string;
} & typeof defaultProps;

const defaultProps = {
  message: 'Heroklock',
};

export const App = (props: Props): JSX.Element => {
  return (
    <div
      style={{
        width: '100vw',
        height: '100vh',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        fontFamily: 'Verdana, sans-serif',
        backgroundColor: 'black',
        color: 'white',
        gap: 20,
      }}
    >
      <h1>
        { props.message }
      </h1>
      <Form
        onSubmit={(url, dur): void => {
          fetch('/add?url=' + url + '&duration=' + dur)
            .then((res) => res.json())
            .then((data: TimeblockType) => {
              // Add new timeblock to container
              fireEvent(Event.NewBlock, data);
            })
            .catch((err) => {
              console.error(err);
              alert('Failed to add new timeblock!');
            });
        }}
      />
      <TimeblockContainer />
      <Footer />
    </div>
  );
}
App.defaultProps = defaultProps;
