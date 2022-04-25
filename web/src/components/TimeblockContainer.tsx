import React from 'react';
import { Event } from '../util';
import { TimeblockType, Timeblock } from './Timeblock';

type Props = {
  interval: number; // In milliseconds, default of 1 minute
} & typeof defaultProps;

const defaultProps = {
  interval: 60 * 1000, // 1 minute
};

type FetchTimeblockResponse = {
  Length: number;
  Blocks: TimeblockType[];
  Iteration: number;
};

export const TimeblockContainer = (props: Props): JSX.Element => {

  const [timeblocks, setTimeblocks] = React.useState<TimeblockType[]>([]);

  const addNewTimeblockListener = (e: CustomEvent) => {
    const newTimeblock : TimeblockType = e.detail;
    console.log(newTimeblock);
    

    setTimeblocks(timeblocks.concat(newTimeblock));
  }

  const fetchTimeblocks = () => {
    fetch('/list')
    .then((res) => res.json())
    .then((data: FetchTimeblockResponse) => {
      setTimeblocks([...data.Blocks]);
    })
    .catch((err) => {
      console.error(err);
      alert('Failed to fetch timeblocks!');
    });
  }

  // On component mount, fetch timeblocks at an interval and add an event
  // handler to handle adding new timeblock
  React.useEffect(() => {

    // Fetch /list for list of timeblocks
    fetchTimeblocks();

    // Fetch /list at an interval of set seconds and minutes
    const intervalId = setInterval(() => { fetchTimeblocks() }, props.interval);

    // Update the timeblocks on add new block
    window.addEventListener(Event.NewBlock, addNewTimeblockListener as EventListener);

    return () => {
      clearInterval(intervalId);
      window.removeEventListener(Event.NewBlock, addNewTimeblockListener as EventListener)
    };
  }, []);

  if (timeblocks.length == 0) {
    return <div />;
  }

  return (
    <div
      style={{
        maxHeight: '50vh',
        overflowY: 'auto',
        overflowX: 'hidden',
        display: 'flex',
        flexDirection: 'column',
        padding: 10,
        border: 'double white',
      }}
    >
      {
        timeblocks.map((data: TimeblockType) => (
          <Timeblock
            key={data.URL}
            URL={data.URL}
            Duration={data.Duration}
            StartTime={data.StartTime}
            Valid={data.Valid}
            delete={(url: string): boolean => {
              // TODO: return false on delete failed
              console.log(url)
              return true;
            }}
          />
        ))
      }
    </div>
  );
}
TimeblockContainer.defaultProps = defaultProps;
