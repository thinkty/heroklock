import React from 'react';

export type TimeblockType = {
  URL: string;
  Duration: number;
  StartTime: string;
  Valid: boolean;
};

type Props = TimeblockType & {
  delete: (url: string) => boolean;
} & typeof defaultProps;

const defaultProps = {};

export const Timeblock = (props: Props): JSX.Element => {
  const startTime = new Date(props.StartTime);
  const endTime = startTime.getTime() + (props.Duration * 1000 * 60 * 30 );
  const timeLeft = endTime - Date.now();

  return (
    <div
      style={{
        width: '100%',
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'space-around',
        alignItems: 'center',
      }}
    >
      <div>
        {
          props.URL
        }
      </div>
      <div>
        {
          Math.floor(timeLeft / (1000 * 60)) + " / " + (props.Duration * 30)
        }
      </div>
    
    </div>
  );
}
Timeblock.defaultProps = defaultProps;
