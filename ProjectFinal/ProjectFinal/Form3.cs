using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;

namespace ProjectFinal
{
    public partial class FrmMainPegawai : Form
    {
        public FrmMainPegawai()
        {
            InitializeComponent();
        }

        private void btnLogout_Click(object sender, EventArgs e)
        {
            FrmLogin frmLogin = new FrmLogin(); 
            frmLogin.Show();
            this.Hide();
        }

        private void btnUpdate_Click(object sender, EventArgs e)
        {

        }
    }
}
