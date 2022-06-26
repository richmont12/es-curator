import React from 'react';
import logo from './logo.svg';
import './Curator.css'
import { RecordOverview } from './UseCases/RecordOverview/RecordOverview';

export class Curator extends React.Component {
    render() {
        return (
            <div className="Curator">
                <header className='Curator-header'>
                    <img src={logo} className="Curator-logo" alt="Curator logo" />
                </header>
                <main className='Curator-main'>
                    <div>
                        <RecordOverview/>
                    </div>
                </main>
            </div>
        );
    }
};
