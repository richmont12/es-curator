import React from 'react';
import ReactDOM from 'react-dom/client';
import { RecordDetailView } from '../../components/RecordDetailView/RecordDetailView';
import { Record } from '../../components/RecordDetailView/Record';
import './RecordOverview.css'

interface IProperties {

}

interface IRecordOverviewState{
    records: Record[]
}

export class RecordOverview extends React.Component<IProperties, IRecordOverviewState> {
    constructor(props: IProperties) {
        super(props);
        this.state = {
            records: [
                new Record(1, "headline1", "desc1", ["link1"]),
                new Record(2, "headline12", "desc12", ["link2"]),
                new Record(3, "headline13", "desc13", ["link3"]),
            ]
        };
    }

    render() {
        const recordDetailViews = this.state.records.map((record: Record) => {
            return (
                <RecordDetailView
                    key={record.Number}
                    record={record}
                />
            );
        })
        return (
            <div>
                <div>{recordDetailViews}</div>
            </div>
        );
    }
};


