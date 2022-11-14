import React from 'react';

import {
  BeakerIcon, CloudIcon, DocumentIcon,
} from '@heroicons/react/24/outline';
import { useNavigate, useLocation, useParams } from 'react-router-dom';
import { faSlack, faGithub } from '@fortawesome/free-brands-svg-icons';
import Navbar from '../components/nav/Navbar';
import Sidebar from '../components/nav/Sidebar';
import { NavLink } from '../components/nav/types';
import { faComponent } from '../utils/utils';

export default function AppContainer(props: {
  children?: React.ReactNode,
}) {
  const navigate = useNavigate();
  const location = useLocation();
  const { id } = useParams();

  const { children } = props;
  const sidebarSections: {
    key: string,
    name: string,
    links: NavLink[]
  }[] = [{
    key: 'base',
    name: '',
    links: [
      {
        title: 'Dashboard',
        icon: BeakerIcon,
        onClick: () => {
          navigate('/dashboard');
        },
        current: location.pathname.startsWith('/dashboard'),
        key: 'dashboard',
      },
      {
        title: 'Data Silos',
        icon: CloudIcon,
        onClick: () => {
          navigate(`/workspaces/${id}/silos`);
        },
        current: location.pathname.startsWith(`/workspaces/${id}/silos`),
        key: 'data_silos',
      },
    ],
  }, {
    key: 'help',
    name: 'Help',
    links: [
      {
        title: 'Documentation',
        icon: DocumentIcon,
        onClick: () => {
          window.open('https://docs.monoid.co', '_blank');
        },
        current: false,
        key: 'docs',
      },
      {
        title: 'Community',
        icon: faComponent(faSlack),
        onClick: () => {
          window.location.href = 'mailto:vignesh@brist.ai?subject=Feature Request';
        },
        current: false,
        key: 'community',
      },
      {
        title: 'Issues/Feature Requests',
        icon: faComponent(faGithub),
        onClick: () => {
          window.location.href = 'mailto:vignesh@brist.ai';
        },
        current: false,
        key: 'issues',
      },
    ],
  }];

  return (
    <div className="flex flex-col h-full bg-gray-100">
      <Navbar
        links={[]}
        dropdownLinks={[]}
        showDropdown={false}
        hiddenLinks={sidebarSections.map((l) => l.links).flat()}
      />
      <div className="flex items-top flex-grow">
        <Sidebar
          sections={sidebarSections}
        />
        <div className="flex-grow bg-gray-100">
          <main>
            <div className="max-w-7xl mx-auto px-4 py-6 sm:px-6 lg:px-8">
              {children}
            </div>
          </main>
        </div>
      </div>
    </div>
  );
}

AppContainer.defaultProps = {
  children: undefined,
};
