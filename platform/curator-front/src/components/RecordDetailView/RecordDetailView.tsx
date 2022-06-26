import React from 'react';
import './RecordDetailView.css'
import { Record } from './Record';

export interface IRecordDetailViewProperties{
    record: Record
}

export class RecordDetailView extends React.Component<IRecordDetailViewProperties> {
    render() {
        const links = this.props.record.Links.map((link: string) => {
            return (
                <a  href={link} 
                    key={link}
                    target="_blank"
                    className='Curator-RecordDetailView-Link'>
                    {link}
                </a>
            )
        });
        return (
            <div className='Curator-RecordDetailView-Container'>
                <div>{this.props.record.Headline}</div>
                <div className='Curator-RecordDetailView-Description'>
                    {this.props.record.Description}
                </div>
                <div className='Curator-RecordDetailView-Links'>
                    {links}
                </div>
            </div>
        );
    }
};
