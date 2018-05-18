package pfappserver::Form::Config::Provisioning::ipsk;

=head1 NAME

pfappserver::Form::Config::Provisioning::ipsk - Web form for IPKS provisioner

=head1 DESCRIPTION

=cut

use HTML::FormHandler::Moose;
extends 'pfappserver::Form::Config::Provisioning';

has_field 'ssid' =>
  (
   type => 'Text',
   label => 'SSID',
  );

has_field 'psk_size' =>
  (
   type => 'PosInteger',
   default => 8,
   label => 'PSK length',
   tags => { after_element => \&help,
             help => 'This is the length of the PSK key you want to generate. The minimum length is eight characters.' },
  );

has_block definition =>
  (
   render_list => [ qw(id description type category ssid oses psk_size) ],
  );


=head1 COPYRIGHT

Copyright (C) 2005-2018 Inverse inc.

=head1 LICENSE

This program is free software; you can redistribute it and/or
modify it under the terms of the GNU General Public License
as published by the Free Software Foundation; either version 2
of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301,
USA.

=cut

__PACKAGE__->meta->make_immutable unless $ENV{"PF_SKIP_MAKE_IMMUTABLE"};
1;
